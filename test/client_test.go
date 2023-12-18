package test

import (
	"log"
	"testing"
	"udp-basic-communication/client"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/until"
)

func TestClient1(t *testing.T) {
	config := client.Config{
		ClientPort: 8080,
		Tag:        "A",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}

func TestClient2(t *testing.T) {
	config := client.Config{
		ClientPort: 8088,
		Tag:        "B",
	}
	testClient := client.NewClient(config)
	testClient.CreateClient()
}

func TestSlice(t *testing.T) {
	p := "{\"type\":3000,\"content\":\"hello\"}"
	var msg = client_model.ReceiveMsg{}
	until.JsonUnmarshal([]byte(p), &msg)
	log.Println(msg.Request)
}
