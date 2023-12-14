package until

import "encoding/json"

func JsonMarshal(data interface{}) []byte {
	conversion, err := json.Marshal(data)
	if err != nil {
		panic("Failed to convert struct to json")
	}
	return conversion
}

func JsonUnmarshal(data []byte, structType any) {
	err := json.Unmarshal(data, structType)
	if err != nil {
		panic("Failed to convert struct to json")
	}
}
