package client_model

import "net"

type ReceiveMsg struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
	Addr    *net.UDPAddr
}
