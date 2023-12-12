package test

import (
	"testing"
	"udp-basic-communication/client"
)

func TestClient1(t *testing.T) {
	config := client.Config{
		ServerIP:   "0.0.0.0",
		ServerPort: 9988,
		ClientPort: 8080,
		Tag:        "A",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}

func TestClient2(t *testing.T) {
	config := client.Config{
		ServerIP:   "0.0.0.0",
		ServerPort: 9988,
		ClientPort: 8088,
		Tag:        "B",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}
