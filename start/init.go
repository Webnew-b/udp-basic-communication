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
		server.GetClientAddr()
		server.GetClientAddr()
		fmt.Println("begin net \n")
		server.ExchangeClientMsg()
		time.Sleep(time.Second * 10)
		server.ResetClientList()
	}
}
