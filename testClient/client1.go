package main

import "udp-basic-communication/client"

func main() {
	config := client.Config{
		ServerIP:   "0.0.0.0",
		ServerPort: 9988,
		ClientPort: 8088,
		Tag:        "B",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}
