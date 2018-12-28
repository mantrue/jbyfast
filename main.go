package main

import (
	"fmt"
	"net/http"
	"penghui.com/config"
	"penghui.com/lib/web"
	"penghui.com/router"
	"time"
)

func main() {
	config.InitConifg()
	//fmt.Println(config.Conf.Runtime.File)
	go func() {
		tick := time.NewTicker(time.Second * 5)
		for range tick.C {
			fmt.Printf("webServer runing....current http request:%v\n", web.HTTP_FAIL_REQUEST_NUM+web.HTTP_RIGHT_REQUEST_NUM)
		}
	}()

	//升级为websocket
	http.ListenAndServe(":80", router.Register())
}
