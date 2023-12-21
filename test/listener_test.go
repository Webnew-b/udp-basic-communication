package test

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"udp-basic-communication/client/listener"
	"udp-basic-communication/enum/listenerStatus"
)

func TestListenerProperties(t *testing.T) {
	name := "TestListener"
	frequency := uint16(100)
	item := listener.NewListener[string](name, frequency)

	assert.Equal(t, name, item.GetName(), "Listener name mismatch")
	assert.Equal(t, frequency, item.GetStopFrequency(), "Listener frequency mismatch")
}

func TestListenerStopListener(t *testing.T) {
	item := listener.NewListener[string]("TestListener", 100)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		item.Run(func(chan string) {})
	}()

	// 停止 Listener
	go item.StopListener()

	// 等待 Listener 完全停止
	wg.Wait()

	assert.Equal(t, listenerStatus.STOP, item.GetStopSymbol(), "Listener did not stop as expected")
}
