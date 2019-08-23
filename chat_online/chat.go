package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// Group struct
type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// User struct
type User struct {
	Id       string          `json:"id"`
	GroupId  string          `json:"group_id"`
	Name     string          `json:"name"`
	Ws       *websocket.Conn `json:"ws"`
	OnlineAt time.Time       `json:"online_at"` //上线时间
	Dt       time.Duration   `json:"dt"`        //在线时长
}

// Messages struct
type Messages struct {
	Id       string  `json:"id"`        //全局唯一id
	Type     int     `json:"type"`      //数据格式类型
	Message  string  `json:"message"`   //消息实体
	UserId   string  `json:"user_id"`   //用户id
	RoomId   string  `json:"room_id"`   //房间号
	CreateAt string  `json:"create_at"` //创建时间
	Unread   uint8   `json:"unread"`    //1未读消息 2已读消息
	Sec      float64 `json:"sec"`       //录音的秒数
	Nickname string  `json:"nickname"`  //昵称
}

// Chat struct
type Chat struct {
	lock        sync.Mutex
	groups      []*Group
	users       []*User
	MessageChan chan Messages
	signalChan  chan string
}

func NewChat() *Chat {
	return &Chat{
		groups:      []*Group{},
		users:       []*User{},
		MessageChan: make(chan Messages, 10),
		signalChan:  make(chan string, 1),
	}
}

// add  group
func (chat *Chat) AddGroup(id string, name string) {
	chat.lock.Lock()
	defer chat.lock.Unlock()

	//check group is exists
	for _, chatGroup := range chat.groups {
		if chatGroup.Id == id {
			return
		}
	}
	group := &Group{
		Id:   id,
		Name: name,
	}

	chat.groups = append(chat.groups, group)
}

// delete group and delete user under group
func (chat *Chat) DeleteGroup(id string) {
	chat.lock.Lock()
	defer chat.lock.Unlock()

	groups := []*Group{}
	for _, chatGroup := range chat.groups {
		if chatGroup.Id == id {
			continue
		}
		groups = append(groups, chatGroup)
	}
	chat.groups = groups

	users := []*User{}
	//clear group user and close ws
	for _, chatUser := range chat.users {
		if chatUser.GroupId == id {
			chatUser.Ws.Close()
			continue
		}
		users = append(users, chatUser)
	}

	chat.users = users
}

// add user
func (chat *Chat) AddUser(id string, groupId string, name string, ws *websocket.Conn) {
	chat.lock.Lock()
	defer chat.lock.Unlock()

	//check group is exists
	groupExist := false
	for _, chatGroup := range chat.groups {
		if chatGroup.Id == groupId {
			groupExist = true
			break
		}
	}
	if !groupExist {
		printError("add user error: group_id is not exists!")
	}

	//check user id name is exists
	for index, chatUser := range chat.users {
		if chatUser.Id == id {
			chat.users = append(chat.users[:index], chat.users[index+1:]...)
		}
	}

	user := &User{
		Id:       id,
		GroupId:  groupId,
		Name:     name,
		Ws:       ws,
		OnlineAt: time.Now(),
	}

	chat.users = append(chat.users, user)
}

// delete user and close user ws
func (chat *Chat) DeleteUser(id string) {
	chat.lock.Lock()
	defer chat.lock.Unlock()

	users := []*User{}
	for _, chatUser := range chat.users {
		if chatUser.Id == id {
			//close ws
			chatUser.Ws.Close()
			chatUser.Dt = time.Since(chatUser.OnlineAt) / 1e9

			//进行数据跟踪
			go httpPostForm(chatUser)

			continue
		}
		users = append(users, chatUser)
	}

	chat.users = users
}

func httpPostForm(chatUser *User) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("http跟踪异常 保存日志服务器出现问题")
			return
		}
	}()

	jsondata, _ := json.Marshal(chatUser)

	request, _ := http.NewRequest("POST", "http://shop.jtypt.com/index.php?s=/wechat/index/chatlog", strings.NewReader(string(jsondata)))
	//post数据并接收http响应
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("http post 数据跟踪异常 日志保存失败", err.Error())
	}
	defer resp.Body.Close()
}

// get user list
func (chat *Chat) Users() []*User {
	return chat.users
}

func (chat *Chat) GetGroupUsers(groupId string) []*User {
	users := []*User{}
	for _, v := range chat.users {
		if v.GroupId == groupId {
			users = append(users, v)
		}
	}
	return users
}

// get group list
func (chat *Chat) Groups() []*Group {
	return chat.groups
}

func printError(message string) {
	fmt.Println(message)
	os.Exit(0)
}
