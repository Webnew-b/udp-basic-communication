package client_model

import (
	"net"
	"udp-basic-communication/enum/msgType"
	"udp-basic-communication/until"
)

type VerifyMsg struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
	Addr    *net.UDPAddr
}

func (this VerifyMsg) ProcessReceiveQueueMsg() {
	//todo 验证消息
	until.PrintMsg(this.Content)
}

func (this VerifyMsg) BuildMsg(content string) []byte {
	msg := VerifyMsg{}
	msg.Type = msgType.CLIENT_VERITY
	msg.Content = content
	buf := until.JsonMarshal(msg)
	return buf
}
