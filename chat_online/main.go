package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}

	chat = NewChat()
)

//初始化配置
func init() {
	initConfig()
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

	log.Println("http server started on : 8087")
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Println("http server error: " + err.Error())
	}
}

// 获取所有用户列表
func userList(w http.ResponseWriter, req *http.Request) {
	users := chat.Users()
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

// ws主任务
func handleConnections(w http.ResponseWriter, req *http.Request) {

	userId := req.FormValue("user_id")
	roomId := req.FormValue("room_id")
	if userId == "" {
		return
	}
	if roomId == "" {
		return
	}
	log.Println("room_id : " + roomId + " user_id: " + userId + " come in!")

	//将 get 请求升级为 websocket
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println("请求出错:" + err.Error())
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
		case 1:
			// 对分组的用户进行广播
			fmt.Println(chat.GetGroupUsers(msg.RoomId))
			fmt.Println("=========================")
			for _, user := range chat.GetGroupUsers(msg.RoomId) {
				ws := user.Ws
				err := ws.WriteJSON(msg)
				if err != nil {
					chat.DeleteUser(user.Id)
					log.Printf("error: %v", err)
				}
			}
		}

	}
}
