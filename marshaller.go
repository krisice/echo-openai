package echoopenai

import "encoding/json"

type Marshaller interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}

type JSONMarshaller struct{}

func (jm *JSONMarshaller) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (jm *JSONMarshaller) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
