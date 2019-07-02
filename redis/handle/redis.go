package handle

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Args struct {
	Id      int
	GetUser string
}

type Reply struct {
	C string
}

type Redis struct{}

var pool *redis.Pool

func init() {
	options := redis.DialPassword("Jingbanyun426!426")
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "118.190.65.33:6372", options)
		},
	}
}

func (t *Redis) GetK(ctx context.Context, args *Args, reply *Reply) error {
	c := pool.Get()
	defer c.Close()

	username, err := redis.String(c.Do("GET", "1"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	reply.C = username

	return nil
}
