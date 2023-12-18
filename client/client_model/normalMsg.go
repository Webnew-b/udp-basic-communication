package client_model

import (
	"udp-basic-communication/enum/msgType"
	"udp-basic-communication/until"
)

type NormalMsg struct {
	*Request
}

func (this NormalMsg) ProcessReceiveQueueMsg() {
	//todo 显示消息
}

func (this NormalMsg) BuildMsg(content string) []byte {
	msg := NormalMsg{}
	msg.Type = msgType.CLIENT_MSG
	msg.Content = content
	buf := until.JsonMarshal(msg)
	return buf
}
