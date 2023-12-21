package client_model

import (
	"udp-basic-communication/enum/msgType"
	"udp-basic-communication/until"
)

type NormalMsg struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
}

func (this NormalMsg) ProcessReceiveQueueMsg() {
	until.PrintMsg(this.Content)
}

func (this NormalMsg) BuildMsg(content string) []byte {
	msg := new(NormalMsg)
	msg.Type = msgType.CLIENT_MSG
	msg.Content = content
	buf := until.JsonMarshal(msg)
	return buf
}
