package main

import (
	"github.com/smallnest/rpcx/server"
	"redis/handle"
)

func main() {
	s := server.NewServer()
	s.Register(new(handle.Redis), "")
	s.Serve("tcp", ":65432")
}
