package message

type Request struct {
	Type    uint16 `json:"type"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Response struct {
	Type    uint16 `json:"type"`
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}
