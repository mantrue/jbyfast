package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://118.190.65.33:1884").SetClientID("123456")
	opts.SetCleanSession(false)
	opts.Username = "chisj"
	opts.Password = "chisj"

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Publish("/test", 2, false, "{name:'oks'}"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Println("=====模拟发送mqtt======")

}
