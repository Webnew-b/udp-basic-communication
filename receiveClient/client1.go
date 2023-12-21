package main

import (
	"fmt"
	"log"
	"udp-basic-communication/client"
)

func main() {
	config := client.Config{
		TargetClientIP:   "127.0.0.1",
		TargetClientPort: 8080,
		ClientPort:       8081,
		Tag:              "B",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
	testClient.SendMsg("how are you josh")
	log.Println("Start Input")
	sth := make([]byte, 18)
	for {
		scanln, _ := fmt.Scanln(&sth)
		log.Println(scanln)
	}
}
