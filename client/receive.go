package client

import (
	"log"
	"net"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/enum/msgType"
	"udp-basic-communication/until"
)

type receive struct {
	msgChannel  chan []byte
	msgQueue    []client_model.NormalMsg
	VerifyQueue []client_model.VerifyMsg
	Connect     *net.UDPConn
}

func (r *receive) receiveHandle() {
	log.Println("Start receive")
	for {
		addr := r.receiveMsg()
		var msg = client_model.ReceiveMsg{}
		bufMsg := <-r.msgChannel
		until.JsonUnmarshal(bufMsg, &msg)
		switch msg.Type {
		case msgType.CLIENT_VERITY:
			verify := client_model.VerifyMsg{
				Request: msg.Request,
				Addr:    addr,
			}
			r.VerifyQueue = append(r.VerifyQueue, verify)
		case msgType.CLIENT_MSG:
			norMsg := client_model.NormalMsg{
				Request: msg.Request,
			}
			r.msgQueue = append(r.msgQueue, norMsg)
		default:
			log.Println("error msg type")
		}
	}
}

func (r *receive) receiveMsg() *net.UDPAddr {
	buf := make([]byte, 256)
	buffLen, addr, err := r.Connect.ReadFromUDP(buf)
	if err != nil {
		log.Panic("Read msg fail", err)
	}
	r.msgChannel <- buf[:buffLen]
	return addr
}
