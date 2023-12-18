package listener

import (
	"log"
	"time"
	"udp-basic-communication/client/client_model"
)

func StartMsgQueueListener[T client_model.Msg](queue client_model.Queue[T]) {
	for {
		if queue.IsQueueEmpty() {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		msg, err := queue.Pop()
		if err != nil {
			log.Println("Error popping from MsgQueue:", err)
			continue
		}
		msg.ProcessQueueMsg()
	}
}
