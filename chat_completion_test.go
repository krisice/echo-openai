package echoopenai

import (
	"os"
	"testing"
)

func TestChatCompletions(t *testing.T) {
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
	}
	res, err := client.CreateChatCompletion(request)
	if err != nil {
		t.Error("test RetrieveModel func failed")
		return
	}
	t.Logf("test RetrieveModel func successed %v", res)
}
