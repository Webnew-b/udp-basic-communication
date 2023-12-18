package client_model

import "net"

type Request struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
}

type ReceiveMsg struct {
	*Request
	Addr *net.UDPAddr
}

type NormalMsg struct {
	*Request
}

type VerifyMsg struct {
	*Request
	Addr *net.UDPAddr
}

func (this NormalMsg) ProcessQueueMsg() {
	//todo 显示消息
}

func (this VerifyMsg) ProcessQueueMsg() {
	//todo 验证消息
}
