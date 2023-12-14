package client

import (
	"log"
	"net"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/enum/msgType"
)

type send struct {
	Connect   *net.UDPConn
	sendQueue []client_model.NormalMsg
}

func (s *send) sendHandle() {
	for {
		// todo 研究切片如何队列化，如何弹出队列元素和加入队列元素
	}
}

func (s *send) pushMsgToQueue(msg client_model.NormalMsg) {
	s.sendQueue = append(s.sendQueue, msg)
}

func (s *send) buildVerifyMsg() client_model.NormalMsg {
	msg := client_model.NormalMsg{}
	msg.Type = msgType.CLIENT_VERITY
	msg.Content = "can you hear me"
	return msg
}

func (s *send) buildNorMsg(content string) client_model.NormalMsg {
	msg := client_model.NormalMsg{}
	msg.Type = msgType.CLIENT_MSG
	msg.Content = content
	return msg
}

func (s *send) sendMsg(msg []byte) {
	_, err := s.Connect.Write(msg)
	if err != nil {
		log.Println("send msg to other client fail", err)
	}
}
