package echoopenai

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestChatCompletionStream(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	message := ChatCompletionMessage{
		Role:    ChatCompletionRoleUser,
		Content: "Test",
	}
	messages := make([]ChatCompletionMessage, 0)
	messages = append(messages, message)
	request := ChatCompletionRequest{
		Model:    string(GPT3Dot5Turbo),
		Messages: messages,
		Stream:   true,
	}
	stream, err := client.CreateChatCompletionStream(request)
	if err != nil {
		t.Error("test CreateChatCompletionStream func failed")
		return
	}

	for {
		content, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				t.Log("stream recv completion")
			} else {
				t.Errorf("stream recv failed %v", err.Error())
			}
			break
		}
		fmt.Println(content)
	}
}
