package client_model

type Queue[T any] interface {
	Push(item T)
	Pop() (T, error)
	Length() int
	IsQueueEmpty() bool
}

type Msg interface {
	BuildMsg(content string) []byte
}

type ReceiveQueueMsg interface {
	ProcessReceiveQueueMsg()
}
