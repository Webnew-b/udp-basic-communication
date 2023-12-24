package client

import (
	"log"
	"net"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/client/listener"
)

type send struct {
	SendConnect     *net.UDPConn
	SendMsgListener *listener.Listener[client_model.SendMsg]
	sendQueue       client_model.Queue[client_model.SendMsg]
}

func (this *send) sendHandle() {
	this.sendQueue = new(client_model.SendQueue)
	this.SendMsgListener = listener.StartMsgQueueListener[client_model.SendMsg](this.sendQueue)
	this.sendMsgByConnect()
}

func (s *send) sendMsgByConnect() {
	go func() {
		for msg := range s.SendMsgListener.Channel {
			_, err := s.SendConnect.Write(msg)
			if err != nil {
				log.Println("send msg to other client fail", err)
			}
		}
	}()
}
