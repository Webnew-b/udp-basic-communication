package client_model

import (
	"errors"
	"sync"
)

type MsgQueue struct {
	Queue []NormalMsg
	mu    sync.Mutex
}

func (this *MsgQueue) Push(item NormalMsg) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.Queue = append(this.Queue, item)
}

func (this *MsgQueue) Pop() (NormalMsg, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.Queue) < 1 {
		return NormalMsg{}, errors.New("queue is empty")
	}
	first := this.Queue[0]
	this.Queue = this.Queue[1:]
	return first, nil
}

func (this *MsgQueue) Length() int {
	return len(this.Queue)
}

func (this *MsgQueue) IsQueueEmpty() bool {
	return this.Length() == 0
}
