package service

import (
	"fmt"
	"log"
	"net"
	"udp-basic-communication/enum/msgStatus"
	"udp-basic-communication/message"
	"udp-basic-communication/until"
)

// Deprecated
type Server struct {
	port       int
	protocol   string
	listener   *net.UDPConn
	clientAddr []*net.UDPAddr
}

// Deprecated
func NewServer(port int, protocol string) *Server {
	server := Server{}
	server.port = port
	server.protocol = protocol
	server.listener, _ = net.ListenUDP(server.protocol, &net.UDPAddr{Port: server.port})
	server.ResetClientList()
	return &server
}

// Deprecated
func (server *Server) CloseServe() {
	err := server.listener.Close()
	if err != nil {
		panic(err)
	}
}

// Deprecated
func (server *Server) SendMsgToClient(msg []byte, addr *net.UDPAddr) {
	_, err := server.listener.WriteToUDP(msg, addr)
	if err != nil {
		msg := fmt.Sprintf("addr:%s:%s;", addr.IP, addr.Port)
		log.Println(msg, "err:", err)
	}
}

// Deprecated
func (server *Server) GetClientAddr() *net.UDPAddr {
	log.Println("getting client")
	bufFormClient := make([]byte, 256)
	n, addr, err := server.listener.ReadFromUDP(bufFormClient)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("read from <%s>:%s\n", addr.String(), bufFormClient[:n])
	log.Println("got client")
	server.clientAddr = append(server.clientAddr, addr)
	return addr
}

// Deprecated
func (server *Server) ReceiveClient(addr *net.UDPAddr) {
	response := message.Response{}
	response.Type = 0
	response.Code = uint16(msgStatus.SUCCESS)
	response.Tag = "server"
	response.Content = "The server has received your message"
	response.Message = msgStatus.MsgStatus.String(msgStatus.SUCCESS)
	responseBuf := until.JsonMarshal(response)
	server.SendMsgToClient(responseBuf, addr)
}

// Deprecated
func (server *Server) ExchangeClientMsg() {
	server.SendMsgToClient([]byte(server.clientAddr[0].String()), server.clientAddr[1])
	server.SendMsgToClient([]byte(server.clientAddr[1].String()), server.clientAddr[0])
}

// Deprecated
func (server *Server) ResetClientList() {
	server.clientAddr = []*net.UDPAddr{}
}
