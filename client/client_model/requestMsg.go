package client_model

type Request struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
}
