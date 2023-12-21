package listener

import (
	"log"
	"sync"
	"time"
	"udp-basic-communication/client/client_model"
)

var (
	mutex      sync.Mutex
	stopCond   = sync.NewCond(&mutex)
	stopChan   = make(chan bool)
	stopSymbol = false
)

func StartReceiveMsgQueueListener[T client_model.ReceiveQueueMsg](queue client_model.Queue[T], isWorking *bool) {
	for {
		mutex.Lock()
		for stopConditionNotMet() { // 这是一个示例条件检查函数
			stopCond.Wait() // 等待停止信号
		}
		mutex.Unlock()

		// 一旦收到停止信号，退出循环
		if stopNeeded() {
			*isWorking = false
			break
		}

		// 检查队列是否为空
		if queue.IsQueueEmpty() {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		*isWorking = true
		msg, err := queue.Pop()
		if err != nil {
			log.Println("Error popping from MsgQueue:", err)
			continue
		}
		msg.ProcessReceiveQueueMsg()

	}
}

func StopAllListeners() {
	mutex.Lock()
	stopSymbol = true
	stopChan <- true
	mutex.Unlock()
	stopCond.Broadcast() // 发送广播通知所有监听器
}

func stopConditionNotMet() bool {
	return !stopSymbol
}

func stopNeeded() bool {
	return stopSymbol
}
