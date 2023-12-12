package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerIP   string
	ServerPort int
	ClientPort int
	Tag        string
}

type Client struct {
	*Config

	serverAddr    *net.UDPAddr
	serverConnect *net.UDPConn

	clientAddr *net.UDPAddr

	otherClientConnect *net.UDPConn
	otherClientAddr    *net.UDPAddr
}

func NewClient(config Config) *Client {
	return &Client{
		Config: &config,
	}
}

func (c *Client) CreateClient() {
	log.Println("client starting")
	c.setClientAddr()
	c.setServerAddr(c.ServerIP, c.ServerPort)
	c.serverConnect = c.connectUDPClient(c.clientAddr, c.serverAddr)
	c.sendMsgToServer("I am ")
	serverMsg := c.readMsg()
	c.otherClientAddr = c.parseAddr(string(serverMsg))
	c.closeServerConnection()
	c.otherClientConnect = c.connectUDPClient(c.clientAddr, c.otherClientAddr)
	c.receiveClientMsg()
	c.sendMsgToOtherClient("test connect")
	c.startInput()
}

func (c *Client) connectUDPClient(from, to *net.UDPAddr) *net.UDPConn {
	log.Println("start connecting")

	conn, err := net.DialUDP("udp", from, to)
	if err != nil {
		log.Fatalln("UDP connect fail", err)
	}
	log.Println("Connect success")
	return conn
}

func (c *Client) sendMsgToOtherClient(msg string) {
	_, err := c.otherClientConnect.Write([]byte(msg))
	if err != nil {
		log.Println("send msg to other client fail", err)
	}
}

func (c *Client) receiveClientMsg() {
	log.Println("Start receive")
	go func() {
		buf := make([]byte, 256)
		for {
			//接受UDP消息打印
			buffLen, _, err := c.otherClientConnect.ReadFromUDP(buf)
			if err != nil {
				log.Println(err)
			}
			if buffLen > 0 {
				fmt.Printf("get msg:%sp2p>", buf[:buffLen])
			}
		}
	}()
}

func (c *Client) startInput() {
	log.Println("Start Input")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p>")
		//读取标准输入，以换行为读取标志
		data, _ := reader.ReadString('\n')
		c.sendMsgToOtherClient(data)
	}

}

func (c *Client) readMsg() []byte {
	buf := make([]byte, 1024*4)
	buffLen, _, err := c.serverConnect.ReadFromUDP(buf)
	if err != nil {
		log.Panic("Parse msg fail", err)
	}
	return buf[:buffLen]
}

func (c *Client) sendMsgToServer(msg string) {
	text := []byte(msg + c.Tag)
	log.Println(text)
	_, err := c.serverConnect.Write(text)
	if err != nil {
		panic(err)
	}
}

func (c *Client) setClientAddr() {
	c.clientAddr = new(net.UDPAddr)
	c.clientAddr.IP = net.ParseIP("127.0.0.1")
	c.clientAddr.Port = c.ClientPort
}

func (c *Client) setServerAddr(ip string, port int) {
	c.serverAddr = new(net.UDPAddr)
	c.serverAddr.IP = net.ParseIP(ip)
	c.serverAddr.Port = port
}

func (c *Client) closeServerConnection() {
	err := c.serverConnect.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Client) parseAddr(addr string) *net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return &net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}
