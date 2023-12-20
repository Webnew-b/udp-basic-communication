package client

import (
	"log"
	"net"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/client/listener"
	"udp-basic-communication/enum/msgType"
	"udp-basic-communication/until"
)

type Receive struct {
	msgChannel                  chan []byte
	ReceiveQueue                client_model.Queue[client_model.NormalMsg]
	VerifyQueue                 client_model.Queue[client_model.VerifyMsg]
	ReceiveConnect              *net.UDPConn
	VerifyQueueListenerWorking  bool
	ReceiveQueueListenerWorking bool
}

func (this *Receive) receiveHandle() {
	this.VerifyQueue = new(client_model.VerifyQueue)
	this.ReceiveQueue = new(client_model.MsgQueue)
	this.ReceiveQueueListenerWorking = false
	this.VerifyQueueListenerWorking = false

	log.Println("Start receive")
	go this.receiveMsg()
	//go listener.StartReceiveMsgQueueListener[client_model.VerifyMsg](this.VerifyQueue, this.VerifyQueueListenerWorking)
	go listener.StartReceiveMsgQueueListener[client_model.NormalMsg](this.ReceiveQueue, &this.ReceiveQueueListenerWorking)
	log.Println("start finish")
}

func (this *Receive) receiveMsg() {
	for {
		buf := make([]byte, 256)
		buffLen, addr, err := this.ReceiveConnect.ReadFromUDP(buf)
		if err != nil {
			log.Panic("Read msg fail:", err)
			return
		}
		var msg = client_model.ReceiveMsg{}
		until.JsonUnmarshal(buf[:buffLen], &msg)
		msg.Addr = addr
		this.putMessageToQueue(&msg)
	}
}

func (this *Receive) putMessageToQueue(msg *client_model.ReceiveMsg) {
	switch msg.Type {
	case msgType.CLIENT_VERITY:
		verify := client_model.VerifyMsg{
			Type:    msg.Type,
			Content: msg.Content,
			Addr:    msg.Addr,
		}
		this.VerifyQueue.Push(verify)
	case msgType.CLIENT_MSG:
		norMsg := client_model.NormalMsg{
			Type:    msg.Type,
			Content: msg.Content,
		}
		this.ReceiveQueue.Push(norMsg)
	default:
		log.Println("error msg type")
	}
}
