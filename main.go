package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
)

func main() {
	server := "115.28.78.221:6379"

	option := redis.DialPassword("Jingbanyun426!426")
	c, err := redis.Dial("tcp", server, option)
	if err != nil {
		log.Println("connect server failed:", err)
		return
	}

	defer c.Close()
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {

			v, err := redis.Int64(c.Do("INCR", "mykey"))
			if err != nil {
				log.Println("INCR failed:", err)
				return
			}
			log.Println("value:", v)
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("end")

}
