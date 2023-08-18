package echoopenai

import "testing"

func TestMashallerUnmashall(t *testing.T) {
	mashaller := &JSONMarshaller{}
	request := ChatCompletionRequest{
		Model:  string(GPT3Dot5Turbo0613),
		N:      1,
		Stream: false,
	}
	_, err := mashaller.Marshal(request)
	if err != nil {
		t.Errorf("test mashaller unmashal func failed %v", err)
	}
	t.Log("test mashaller unmashal func successed")
}
