package model

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
)

type Args struct {
	Id      int
	GetUser string
}

type Reply struct {
	C string
}

func UserKeyId() (string, error) {
	fmt.Println("rpc client run.....................")

	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:65432", "")

	xclient := client.NewXClient("Redis", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Args{}
	reply := &Reply{}

	fmt.Println("rpc call start..........................")
	err := xclient.Call(context.Background(), "GetK", args, reply)
	fmt.Println("rpc call end..........................")

	if err != nil {
		panic(err)
	}

	return reply.C, nil
}
