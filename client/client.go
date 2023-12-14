package client

import (
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
	*receive
	*send

	clientAddr *net.UDPAddr

	clientToken string

	targetClientConnect *net.UDPConn
	targetClientAddr    *net.UDPAddr
}

func NewClient(config Config) *Client {
	client := &Client{
		Config: &config,
	}
	client.setClientAddr()
	return client
}

func (c *Client) CreateClient() {
	log.Println("client starting")
	c.receiveHandle()
	c.sendHandle()
}

func (c *Client) connectTargetClient() {
	// todo 连接逻辑
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

//todo 判断消息类型
//todo 验证消息
//todo 显示消息

func (c *Client) startInput() {
	// todo 这个输入拓展成所有输入，不仅仅只是发消息
	log.Println("Start Input")
}

/*func (c *Client) sendMsgToServer(msg string) {
	text := message.Request{}

	text.Type = msgType.SERVER_VERITY
	text.Tag = c.Tag
	text.Content = msg
	buf := until.JsonMarshal(text)

	_, err := c.serverConnect.Write(buf)
	if err != nil {
		panic(err)
	}
}*/

func (c *Client) setClientAddr() {
	c.clientAddr = new(net.UDPAddr)
	c.clientAddr.IP = net.ParseIP("127.0.0.1")
	c.clientAddr.Port = c.ClientPort
}

func (c *Client) SetTargetClientAddr(ip string, port int) {
	c.targetClientAddr = new(net.UDPAddr)
	c.targetClientAddr.IP = net.ParseIP(ip)
	c.targetClientAddr.Port = port
}

func (c *Client) CloseConnection() {
	err := c.targetClientConnect.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

/*func (c *Client) parseAddr(addr string) *net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return &net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}*/
