package listener

import (
	"log"
	"time"
	"udp-basic-communication/client/client_model"
)

func StartReceiveMsgQueueListener[T any](queue client_model.Queue[T]) *Listener[T] {

	queueListener := NewListener[T]("ReceiveQueue", 100)

	go queueListener.Run(
		func(channel chan T) {
			if queue.IsQueueEmpty() {
				time.Sleep(time.Millisecond * 100)
				return
			}
			msg, err := queue.Pop()
			channel <- msg
			if err != nil {
				log.Println("Error popping from MsgQueue:", err)
				return
			}
		},
	)

	return queueListener
}
