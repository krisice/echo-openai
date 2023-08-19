package echoopenai

import (
	"context"
	"errors"
	"net/http"
)

var ErrUnexpectedEmbeddingRequestType = errors.New("error: unexpected input field in request")

type EmbeddingRequest struct {
	Model OpenAIModel `json:"model"`
	Input any         `json:"input"`
	User  string      `json:"user"`
}

type EmbeddingData struct {
	Index     int       `json:"index"`
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
}

type EmbeddingResponse struct {
	Object string          `json:"object"`
	Model  string          `json:"model"`
	Data   []EmbeddingData `json:"data"`
	Usage  TokenUsage      `json:"usage"`
}

func (c *Client) CreateEmbeddings(request EmbeddingRequest) (response EmbeddingResponse, err error) {
	return c.CreateEmbeddingsWithContext(context.Background(), request)
}

func (c *Client) CreateEmbeddingsWithContext(ctx context.Context, request EmbeddingRequest) (response EmbeddingResponse, err error) {
	urlSuffix := "embeddings"

	if !isString(request.Input) && !isSlice(request.Input) {
		err = ErrUnexpectedEmbeddingRequestType
		return
	}

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), request)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}
