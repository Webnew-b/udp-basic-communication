package client_model

import (
	"errors"
	"sync"
)

type SendMsg []byte

type SendQueue struct {
	Queue []SendMsg
	mu    sync.Mutex
}

func (this *SendQueue) Push(item SendMsg) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Queue = append(this.Queue, item)
}

func (this *SendQueue) Pop() (SendMsg, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.Queue) < 1 {
		return SendMsg{}, errors.New("queue is empty")
	}
	first := this.Queue[0]
	this.Queue = this.Queue[1:]
	return first, nil
}

func (this *SendQueue) Length() int {
	return len(this.Queue)
}

func (this *SendQueue) IsQueueEmpty() bool {
	return this.Length() == 0
}
