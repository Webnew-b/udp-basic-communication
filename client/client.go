package client

import (
	"fmt"
	"log"
	"net"
	"time"
	"udp-basic-communication/client/client_model"
	"udp-basic-communication/enum/msgType"
)

type Config struct {
	TargetClientIP   string
	TargetClientPort int
	ClientPort       int
	Tag              string
}

type Client struct {
	*Config
	*Receive
	*send

	clientAddr *net.UDPAddr

	clientToken string

	targetClientConnect *net.UDPConn
	targetClientAddr    *net.UDPAddr
}

func NewClient(config Config) *Client {
	client := &Client{
		Receive: new(Receive),
		send:    new(send),
		Config:  &config,
	}
	client.setTargetClientAddr()
	client.setClientAddr()
	return client
}

func (c *Client) CreateClient() {
	log.Println("client starting")
	c.setReceiveAndSendConn(c.targetClientAddr)
	c.receiveHandle()
	c.sendHandle()
	//c.startInput()
}

func (c *Client) setReceiveAndSendConn(addr *net.UDPAddr) {
	connect := c.createUDPConnection(c.clientAddr, addr)
	c.ReceiveConnect = connect
	c.SendConnect = connect
}

func (c *Client) connectTargetClient() {
	msg := new(client_model.VerifyMsg)
	buf := msg.BuildVerifyMsg("request Connecting By " + msg.Addr.String())
	c.sendQueue.Push(buf)
	log.Println("verify message be sent")
	c.handleVerifyReply()
}

func (c *Client) agreeConnect() {
	msg := new(client_model.VerifyMsg)
	buf := msg.BuildAgreeMsg()
	c.sendQueue.Push(buf)
}

func (c *Client) rejectConnect() {
	msg := new(client_model.VerifyMsg)
	buf := msg.BuildRejectMsg()
	c.sendQueue.Push(buf)
}

func (c *Client) handleVerifyReply() {
	go func() {
		select {
		case verifyMsg := <-c.VerifyMsgListener.Channel:
			c.determiningMsgType(verifyMsg)
		case <-time.After(5 * time.Second):
			log.Println("connection is time out")
		}
	}()
}

func (c *Client) determiningMsgType(msg client_model.VerifyMsg) {
	switch msg.Type {
	case msgType.CLIENT_AGREE:
		c.setReceiveAndSendConn(msg.Addr)
	case msgType.CLIENT_REJECT:
		addr := msg.Addr.String()
		printStr := fmt.Sprintf("client(%s) reject connection", addr)
		log.Println(printStr)
	}
}

func (c *Client) createUDPConnection(from, to *net.UDPAddr) *net.UDPConn {
	log.Println("start connecting")
	conn, err := net.DialUDP("udp", from, to)
	if err != nil {
		log.Fatalln("UDP connect fail", err)
	}
	log.Println("Connect success")
	return conn
}

func (c *Client) startInput() {
	// todo 这个输入拓展成所有输入，不仅仅只是发消息
}

func (c *Client) setClientAddr() {
	c.clientAddr = new(net.UDPAddr)
	c.clientAddr.IP = net.ParseIP("127.0.0.1")
	c.clientAddr.Port = c.ClientPort
}

func (c *Client) setTargetClientAddr() {
	c.targetClientAddr = new(net.UDPAddr)
	c.targetClientAddr.IP = net.ParseIP(c.TargetClientIP)
	c.targetClientAddr.Port = c.TargetClientPort
}

func (c *Client) CloseConnection() {
	err := c.targetClientConnect.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Client) SendMsg(content string) {
	msg := new(client_model.NormalMsg)
	buf := msg.BuildMsg(content)
	c.sendQueue.Push(buf)
}

func (c *Client) handleMsg(msgChannel chan client_model.NormalMsg) {
	for msg := range msgChannel {
		log.Println(msg.Content)
	}
}
