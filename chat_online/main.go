package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域(一般来讲,websocket都是独立部署的)
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	chat = NewChat()

	pool *redis.Pool
)

//初始化配置
func init() {

	options := redis.DialPassword("Jingbanyun426!426")
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "115.28.78.221:6379", options)
		},
	}

}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	// 用户列表
	http.HandleFunc("/users", userList)
	// 分组列表
	http.HandleFunc("/groups", groupList)
	//根据分组id分组在线列表
	http.HandleFunc("/gusers", groupIdList)
	// ws主信息
	http.HandleFunc("/ws", handleConnections)

	// 开启协程读取ws信息
	go handleMessages()

	//开启心跳
	go pingClient()

	log.Println("http server started on : 8087")
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Println("http server error: " + err.Error())
	}
}

// 获取所有用户列表
func userList(w http.ResponseWriter, req *http.Request) {
	users := chat.Users()

	for k, v := range users {
		users[k].Dt = time.Since(v.OnlineAt) / 1e9
	}

	v, _ := json.Marshal(users)
	w.Write(v)
}

//根据分组id获取用户列表

func groupIdList(w http.ResponseWriter, req *http.Request) {
	group_id := req.FormValue("group_id")
	users := chat.GetGroupUsers(group_id)
	v, _ := json.Marshal(users)
	w.Write(v)
}

// 获取所有分组
func groupList(w http.ResponseWriter, req *http.Request) {
	groups := chat.Groups()
	v, _ := json.Marshal(groups)
	w.Write(v)
}

//ping client
func pingClient() {
	tk := time.NewTicker(time.Second * 5)
	for {
		<-tk.C
		if len(chat.Users()) > 0 {
			for _, user := range chat.Users() {
				ws := user.Ws
				err := ws.WriteMessage(websocket.TextMessage, []byte("ping"))
				if err != nil {
					chat.DeleteUser(user.Id)
					log.Printf("极端情况离线error: %v", err)
				}
			}
		}
	}
}

// ws主任务
func handleConnections(w http.ResponseWriter, req *http.Request) {
	//fmt.Println("打印头部信息", req.Header)

	userId := req.FormValue("user_id")
	roomId := req.FormValue("room_id")
	if userId == "" {
		log.Println("用户ID为空")
		return
	}
	if roomId == "" {
		log.Println("房间ID为空")
		return
	}
	log.Println("room_id : " + roomId + " user_id: " + userId + " come in!")

	//将 get 请求升级为 websocket
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("请求出错:" + err.Error())
		return
	}
	defer ws.Close()

	// 添加到分组
	chat.AddGroup(roomId, "default_"+roomId)
	// 添加用户
	chat.AddUser(userId, roomId, "user_"+userId, ws)

	for {
		// 读取前台的信息
		var msg Messages
		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Println(err.Error())
			log.Println("room_id : " + roomId + " user_id: " + userId + " exit!")
			chat.DeleteUser(userId)
			break
		}

		// 发送信息到队列
		chat.MessageChan <- msg
	}
}

// 从队列中读取信息
func handleMessages() {
	for {
		// 定义读取信息的格式
		var msg Messages
		msg = <-chat.MessageChan

		switch msg.Type {
		case 1: //全局进行房间广播

			//开启一个新的协程进行入库操作 存入redis
			msg.CreateAt = time.Now().Format("2006-01-02 15:04:05")
			id := xid.New()
			msg.Id = id.String() //生成全局唯一的id
			msg.Unread = 1       //发过来的消息默认都是未读状态

			go saveChatRedis(msg)
			// 对分组的用户进行广播

			for _, user := range chat.GetGroupUsers(msg.RoomId) {
				ws := user.Ws
				err := ws.WriteJSON(msg)
				if err != nil {
					chat.DeleteUser(user.Id)
					log.Printf("error: %v", err)
					log.Println("用户断开连接" + err.Error())
				}
			}

		case 2: //进行心跳检测
			fmt.Println("进行心跳检测")
		}

	}
}

//把聊天数据入库到redis中
func saveChatRedis(msg Messages) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("redis跟踪异常 数据无法保存")
			return
		}
	}()

	msgJ, _ := json.Marshal(msg)
	c := pool.Get()
	_, err := c.Do("zadd", "chat:"+msg.RoomId, time.Now().UnixNano()/1e6, msgJ)
	if err != nil {
		log.Println("redis保存聊天失败" + err.Error())
		return
	}

	_, err = c.Do("lpush", "push:chat:"+msg.RoomId, msgJ)
	if err != nil {
		log.Println("redis保存推送失败" + err.Error())
		return
	}
	return
}
