package main

import (
	"log"
	"time"
	"udp-basic-communication/client"
)

func main() {
	config := client.Config{
		TargetClientIP:   "127.0.0.1",
		TargetClientPort: 8081,
		ClientPort:       8080,
		Tag:              "A",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
	time.Sleep(time.Second * 5)
	log.Println("send Msg")
	testClient.SendMsg("how are you Tom")
	time.Sleep(time.Second * 5)
}
