package basic_date

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	fmt.Println("staring the server...")
	listen, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}

		go doReadBufio(conn)
	}
}

func doReadBufio(conn net.Conn) {
	for {
		buf := make([]byte, 10)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println("Received data:", time.Now(), string(buf))
	}
}
