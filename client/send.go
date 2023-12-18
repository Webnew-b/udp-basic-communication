package client

import (
	"log"
	"net"
	"time"
	"udp-basic-communication/client/client_model"
)

type send struct {
	SendConnect *net.UDPConn
	sendQueue   client_model.Queue[client_model.SendMsg]
}

func (this *send) sendHandle() {
	this.sendQueue = new(client_model.SendQueue)
	go this.startMsgQueueListener()
}

func (this *send) sendMsgByConnect(msg []byte) {
	_, err := this.SendConnect.Write(msg)
	if err != nil {
		log.Println("send msg to other client fail", err)
	}
}

func (this *send) startMsgQueueListener() {
	for {
		if this.sendQueue.IsQueueEmpty() {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		msg, err := this.sendQueue.Pop()
		if err != nil {
			log.Println("Error popping from MsgQueue:", err)
			continue
		}
		this.sendMsgByConnect(msg)
	}
}
