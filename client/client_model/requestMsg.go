package client_model

import "net"

type Request struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
}

type ReceiveMsg struct {
	*Request
}

type NormalMsg struct {
	*Request
}

type VerifyMsg struct {
	*Request
	Addr *net.UDPAddr
}
