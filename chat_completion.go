package echoopenai

import (
	"context"
	"net/http"
)

type ChatCompletionRole string

const (
	ChatCompletionRoleAssistant ChatCompletionRole = "assistant"
	ChatCompletionRoleUser      ChatCompletionRole = "user"
	ChatCompletionRoleSystem    ChatCompletionRole = "system"
)

type ChatCompletionMessage struct {
	Role    ChatCompletionRole `json:"role"`
	Content string             `json:"content"`
	Name    string             `json:"name,omitempty"`
}

type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	User             string                  `json:"user,omitempty"`
	N                int32                   `json:"n,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
	MaxTokens        int32                   `json:"max_tokens,omitempty"`
	Temperature      float32                 `json:"temperature,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
}

type ChatCompletionChoices struct {
	Index        int32                 `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
}

type ChatCompletionResponse struct {
	ID      string                  `json:"id"`
	Object  string                  `json:"object"`
	Created int64                   `json:"created"`
	Model   string                  `json:"model"`
	Choices []ChatCompletionChoices `json:"choices"`
	Usage   TokenUsage              `json:"usage"`
}

func (c *Client) CreateChatCompletion(request ChatCompletionRequest) (ChatCompletionResponse, error) {
	return c.CreateChatCompletionWithContext(context.Background(), request)
}

func (c *Client) CreateChatCompletionWithContext(ctx context.Context, request ChatCompletionRequest) (response ChatCompletionResponse, err error) {
	urlSuffix := "chat/completions"
	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), request)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	c.sendRequestWithContext(ctx, req, &response)
	return
}
