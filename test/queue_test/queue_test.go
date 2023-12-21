package queue_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"udp-basic-communication/client/client_model"
)

func TestMsgQueue(t *testing.T) {
	mq := &client_model.MsgQueue{Queue: []client_model.NormalMsg{}}
	testItems := []client_model.NormalMsg{client_model.NormalMsg{}, client_model.NormalMsg{}}
	StartTestQueue[client_model.NormalMsg](t, mq, testItems)
}

func TestVerifyQueue(t *testing.T) {
	mq := &client_model.VerifyQueue{Queue: []client_model.VerifyMsg{}}
	testItems := []client_model.VerifyMsg{client_model.VerifyMsg{}, client_model.VerifyMsg{}}
	StartTestQueue[client_model.VerifyMsg](t, mq, testItems)
}

func TestSendQueue(t *testing.T) {
	mq := &client_model.SendQueue{Queue: []client_model.SendMsg{}}
	testItems := []client_model.SendMsg{client_model.SendMsg{}, client_model.SendMsg{}}
	StartTestQueue[client_model.SendMsg](t, mq, testItems)
}

func StartTestQueue[T any](t *testing.T, q client_model.Queue[T], testItems []T) {
	assert.Equal(t, 0, q.Length())
	assert.True(t, q.IsQueueEmpty())

	for _, item := range testItems {
		q.Push(item)
	}

	assert.Equal(t, len(testItems), q.Length())

	for _, expectedItem := range testItems {
		item, err := q.Pop()
		assert.NoError(t, err)
		assert.Equal(t, expectedItem, item)
	}

	assert.Equal(t, 0, q.Length())
	assert.True(t, q.IsQueueEmpty())

	_, err := q.Pop()
	assert.Error(t, err) // 当队列为空时，Pop 应该返回错误
}
