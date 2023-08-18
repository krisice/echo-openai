package echoopenai

import "testing"

func TestChatCompletions(t *testing.T) {
	client := NewClient("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")
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
