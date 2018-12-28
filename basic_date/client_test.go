package basic_date

import (
	"net"
	"testing"
)

func TestSocketSend(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		t.Error("error")
		return
	}
	s, err := conn.Write([]byte("123456"))
	if err != nil {
		t.Error(err, s)
	}
	defer conn.Close()
}
