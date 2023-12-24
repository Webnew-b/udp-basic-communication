package listener

import (
	"log"
	"time"
	"udp-basic-communication/client/client_model"
)

func StartMsgQueueListener[T any](queue client_model.Queue[T]) *Listener[T] {

	queueListener, err := NewListener[T]("ReceiveQueueListener", 100)

	if err != nil {
		panic(err)
	}

	err = queueListener.Run(
		func(channel *chan T) {
			if queue.IsQueueEmpty() {

				time.Sleep(time.Millisecond * 100)
				return
			}
			msg, err := queue.Pop()
			*channel <- msg
			if err != nil {
				log.Println("Error popping from MsgQueue:", err)
				return
			}
		},
	)
	if err != nil {
		panic(err)
	}

	return queueListener
}
