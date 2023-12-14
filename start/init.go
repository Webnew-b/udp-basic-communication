package start

import (
	"fmt"
	"log"
	"time"
	"udp-basic-communication/service"
)

func StartServe() {
	var server = service.NewServer(9988, "udp")
	//defer server.CloseServe()
	for {
		log.Println("start Server")
		firstClientAddr := server.GetClientAddr()
		server.ReceiveClient(firstClientAddr)
		secondClientAddr := server.GetClientAddr()
		server.ReceiveClient(secondClientAddr)
		fmt.Println("begin net \n")
		server.ExchangeClientMsg()
		time.Sleep(time.Second * 10)
		server.ResetClientList()
	}
}
