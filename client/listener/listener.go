package listener

import (
	"sync"
	"time"
	"udp-basic-communication/enum/listenerStatus"
)

/*var (
	mutex      sync.Mutex
	stopCond   = sync.NewCond(&mutex)
	stopChan   = make(chan bool)
	stopSymbol = false
)*/

type Listener[T any] struct {
	name       string
	stopSymbol listenerStatus.Status
	frequency  uint16

	mutex    sync.Mutex
	stopCond *sync.Cond
	channel  chan T
}

func NewListener[T any](name string, frequency uint16) *Listener[T] {
	listener := Listener[T]{
		name:       name,
		stopSymbol: listenerStatus.STOP,
		frequency:  frequency,
	}
	listener.mutex.Lock()
	listener.stopCond = sync.NewCond(&listener.mutex)
	return &listener
}

func (this *Listener[T]) Run(action func(channel chan T)) {
	for {
		isStop := this.HandleStatusChange()
		if isStop {
			break
		}
		action(this.channel)
	}
}

func (this *Listener[T]) HandleStatusChange() bool {
	for this.stopConditionNotMet() {
		this.stopCond.Wait()
	}
	if this.stopNeeded() {
		return true
	}
	return false
}

func (this *Listener[T]) StopListener() {
	this.stopSymbol = listenerStatus.READY_STOP
	second := int(this.frequency) + 500
	this.stopCond.Broadcast()
	time.Sleep(time.Duration(second) * time.Millisecond)
	this.stopSymbol = listenerStatus.STOP
	time.Sleep(time.Duration(second) * time.Millisecond)
}

func (this *Listener[T]) stopConditionNotMet() bool {
	return this.stopSymbol == listenerStatus.READY_STOP
}

func (this *Listener[T]) stopNeeded() bool {
	return this.stopSymbol == listenerStatus.STOP
}

func (this *Listener[T]) GetName() string {
	return this.name
}

func (this *Listener[T]) GetStopSymbol() listenerStatus.Status {
	return this.stopSymbol
}

func (this *Listener[T]) GetStopFrequency() uint16 {
	return this.frequency
}
