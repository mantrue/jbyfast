package main

import (
	"fmt"
	"net/http"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	// 创建记录日志的文件
	f, _ := os.Create("request.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
