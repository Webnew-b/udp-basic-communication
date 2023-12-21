package listener_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"udp-basic-communication/client/listener"
	"udp-basic-communication/enum/listenerStatus"
)

// TestNewListener 测试 listener.NewListener 方法的功能和错误处理。
func TestNewListener(t *testing.T) {
	// 测试正常情况
	testListener, err := listener.NewListener[string]("TestListener", 100)
	assert.Nil(t, err)
	assert.NotNil(t, testListener)

	// 测试错误情况（例如频率为零）
	testListener, err = listener.NewListener[string]("TestListener", 0)
	assert.NotNil(t, err)
	assert.Nil(t, testListener)
}

// TestRun 测试 listener.Listener 的 Run 方法。
func TestRun(t *testing.T) {
	testListener, _ := listener.NewListener[string]("TestListener", 100)

	// 第一次运行
	err := testListener.Run(func(ch chan string) {})
	assert.Nil(t, err)

	// 测试重复运行
	err = testListener.Run(func(ch chan string) {})
	assert.NotNil(t, err)

	// 测试是否正在运行
	assert.True(t, testListener.IsWorking())

	// 停止 Listener 以允许其他测试进行
	testListener.StopListener()
	time.Sleep(200 * time.Millisecond)
}

// TestStopListener 测试 listener.Listener 的 StopListener 方法。
func TestStopListener(t *testing.T) {
	testListener, _ := listener.NewListener[string]("TestListener", 100)

	// 启动 Listener
	err := testListener.Run(func(ch chan string) {})
	assert.NoError(t, err)
	time.Sleep(100 * time.Millisecond)

	// 停止 Listener
	testListener.StopListener()

	// 等待足够的时间以确保监听器已停止
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, listenerStatus.STOP, testListener.GetStopSymbol())
	assert.False(t, testListener.IsWorking())
}

// TestGetName 测试 listener.Listener 的 GetName 方法。
func TestGetName(t *testing.T) {
	testListener, _ := listener.NewListener[string]("TestListener", 100)
	assert.Equal(t, "TestListener", testListener.GetName())
}

// TestGetStopSymbol 测试 listener.Listener 的 GetStopSymbol 方法。
func TestGetStopSymbol(t *testing.T) {
	testListener, _ := listener.NewListener[string]("TestListener", 100)
	assert.Equal(t, listenerStatus.STOP, testListener.GetStopSymbol()) // 初始状态应为 STOP
}

// TestGetStopFrequency 测试 listener.Listener 的 GetStopFrequency 方法。
func TestGetStopFrequency(t *testing.T) {
	testListener, _ := listener.NewListener[string]("TestListener", 100)
	assert.Equal(t, uint16(100), testListener.GetStopFrequency())
}
