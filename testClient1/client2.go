package main

import "udp-basic-communication/client"

func main() {
	config := client.Config{
		TargetClientIP:   "127.0.0.1",
		TargetClientPort: 8081,
		ClientPort:       8080,
		Tag:              "A",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}
