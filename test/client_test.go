package test

import (
	"log"
	"sync"
	"testing"
	"time"
)

var (
	mutex      sync.Mutex
	stopCond   = sync.NewCond(&mutex)
	stopChan   = make(chan bool)
	stopSignal bool
)

func TestClient1(t *testing.T) {
	go putSymbolToChan()
	go func() {
		for {
			mutex.Lock()
			for stopConditionNotMet() {
				stopCond.Wait() // 等待停止信号
			}
			mutex.Unlock()

			// 一旦收到停止信号，退出循环
			if stopNeeded() {
				break
			}

			// 监听器的正常操作
			if false {
				time.Sleep(time.Millisecond * 500)
				log.Println(123456)
				continue
			}
		}
	}()
}

func TestClient2(t *testing.T) {
	log.Println(!stopSignal)
}

func TestSlice(t *testing.T) {

}

func putSymbolToChan() {
	time.Sleep(time.Second * 50)
	StopAllListeners()
}

func StopAllListeners() {
	mutex.Lock()
	stopSignal = true
	stopChan <- true
	mutex.Unlock()
	stopCond.Broadcast() // 发送广播通知所有监听器
}

// stopConditionNotMet 检查是否收到停止信号的逻辑
func stopConditionNotMet() bool {
	return !stopSignal
}

// stopNeeded 检查是否需要停止监听器的逻辑
func stopNeeded() bool {
	return stopSignal
}
