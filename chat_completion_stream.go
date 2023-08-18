package echoopenai

import (
	"bufio"
	"context"
	"errors"
	"net/http"
)

var (
	ErrTooManyEmptyStreamMessages = errors.New("stream has sent too many empty messages")
)

type ChatCompletionStream struct {
	*streamReader[ChatCompletionStreamResponse]
}

type ChatCompletionStreamDelta struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionStreamChoices struct {
	Index        int32                     `json:"index"`
	Delta        ChatCompletionStreamDelta `json:"delta"`
	FinishReason string                    `json:"finish_reason"`
}

type ChatCompletionStreamResponse struct {
	ID      string                        `json:"id"`
	Object  string                        `json:"object"`
	Created int64                         `json:"created"`
	Model   string                        `json:"model"`
	Choices []ChatCompletionStreamChoices `json:"choices"`
}

func (c *Client) CreateChatCompletionStream(req ChatCompletionRequest) (*ChatCompletionStream, error) {
	return c.CreateChatCompletionStreamWithContext(context.Background(), req)
}

func (c *Client) CreateChatCompletionStreamWithContext(ctx context.Context, request ChatCompletionRequest) (stream *ChatCompletionStream, err error) {
	request.Stream = true
	urlSuffix := "chat/completions"

	req, err := c.createStreamRequestWithContext(ctx, http.MethodPost, urlSuffix, request)
	if err != nil {
		return
	}

	res, err := c.client.Do(req)
	if err != nil {
		return
	}

	if isFailureStatusCode(res) {
		return nil, c.handleErrorResp(res)
	}

	stream = &ChatCompletionStream{
		streamReader: &streamReader[ChatCompletionStreamResponse]{
			emptyMessagesLimit: defaultEmptyMessagesLimit,
			scanner:            bufio.NewScanner(res.Body),
			response:           res,
			errAccumulator:     NewErrorAccumulator(),
			marshaler:          &JSONMarshaller{},
		},
	}
	return
}
