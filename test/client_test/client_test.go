package client_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)
	go func() {
		time.Sleep(8 * time.Second)
		ch <- "123"
	}()
	go func() {
		select {
		case msg := <-ch:
			log.Println(msg)
			assert.True(t, false)
		case <-time.After(5 * time.Second):
			log.Println("yes")
			assert.True(t, true)
		}
	}()
	time.Sleep(15 * time.Second)
}
