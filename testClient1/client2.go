package main

import "Intranet-text-transmission/client"

func main() {
	config := client.Config{
		ServerIP:   "0.0.0.0",
		ServerPort: 9988,
		ClientPort: 8080,
		Tag:        "A",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}
