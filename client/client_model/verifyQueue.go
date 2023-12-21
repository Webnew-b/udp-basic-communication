package client_model

import (
	"errors"
	"sync"
)

type VerifyQueue struct {
	Queue []VerifyMsg
	mu    sync.Mutex
}

func (this *VerifyQueue) Push(item VerifyMsg) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Queue = append(this.Queue, item)
}

func (this *VerifyQueue) Pop() (VerifyMsg, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.Queue) < 1 {
		return VerifyMsg{}, errors.New("queue is empty")
	}
	first := this.Queue[0]
	this.Queue = this.Queue[1:]
	return first, nil
}

func (this *VerifyQueue) Length() int {
	return len(this.Queue)
}

func (this *VerifyQueue) IsQueueEmpty() bool {
	return this.Length() == 0
}
