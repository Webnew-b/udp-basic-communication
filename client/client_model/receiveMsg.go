package client_model

import "net"

type ReceiveMsg struct {
	*Request
	Addr *net.UDPAddr
}
