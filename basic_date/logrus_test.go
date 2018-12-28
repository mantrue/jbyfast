package basic_date

import (
	"crypto/md5"
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.Info("hello, world.")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

func BenchmarkLorus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// logrus.Info("hello, world.")
		fmt.Println("hello world")
	}
}

func TestMd5(t *testing.T) {
	md := md5.Sum([]byte("hello"))
	fmt.Printf("%x\n", md)
}
