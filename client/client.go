package client

import (
	"fmt"
	"log"
	"net"
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
	connect := c.createUDPConnection(c.clientAddr, c.targetClientAddr)
	c.ReceiveConnect = connect
	c.SendConnect = connect
	c.receiveHandle()
	sth := make([]byte, 18)
	for {
		scanln, _ := fmt.Scanln(&sth)
		log.Panicln(scanln)
	}
}

func (c *Client) connectTargetClient() {
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
	log.Println("Start Input")
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
