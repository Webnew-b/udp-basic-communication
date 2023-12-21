package listener

import (
	"fmt"
	"sync"
	"time"
	"udp-basic-communication/enum/listenerStatus"
)

type Listener[T any] struct {
	name       string
	stopSymbol listenerStatus.Status
	frequency  uint16

	mutex     sync.Mutex
	stopCond  *sync.Cond
	Channel   chan T
	isWorking bool
}

func NewListener[T any](name string, frequency uint16) (*Listener[T], error) {
	if frequency == 0 {
		return nil, fmt.Errorf("frequency cannot be zero")
	}
	listener := Listener[T]{
		name:       name,
		stopSymbol: listenerStatus.STOP,
		frequency:  frequency,
	}
	listener.mutex.Lock()
	listener.stopCond = sync.NewCond(&listener.mutex)
	return &listener, nil
}

func (l *Listener[T]) Run(action func(channel chan T)) error {
	if l.isWorking {
		return fmt.Errorf("listener is already running")
	}
	l.isWorking = true
	go func() {
		l.stopSymbol = listenerStatus.START
		for {
			isStop := l.HandleStatusChange()
			if isStop {
				l.isWorking = false
				break
			}
			action(l.Channel)
		}
	}()
	return nil
}

func (l *Listener[T]) HandleStatusChange() bool {
	for l.isReadyToStop() {
		l.stopCond.Wait()
	}
	if l.stopNeeded() {
		return true
	}
	return false
}

func (l *Listener[T]) StopListener() {
	go func() {
		l.stopSymbol = listenerStatus.READY_STOP
		second := int(l.frequency) + 500
		time.Sleep(time.Duration(second) * time.Millisecond)
		l.stopSymbol = listenerStatus.STOP
		l.stopCond.Broadcast()
		time.Sleep(time.Duration(second) * time.Millisecond)
	}()
}

func (l *Listener[T]) isReadyToStop() bool {
	return l.stopSymbol == listenerStatus.READY_STOP
}

func (l *Listener[T]) stopNeeded() bool {
	return l.stopSymbol == listenerStatus.STOP
}

func (l *Listener[T]) GetName() string {
	return l.name
}

func (l *Listener[T]) GetStopSymbol() listenerStatus.Status {
	return l.stopSymbol
}

func (l *Listener[T]) GetStopFrequency() uint16 {
	return l.frequency
}

func (l *Listener[T]) IsWorking() bool {
	return l.isWorking
}
