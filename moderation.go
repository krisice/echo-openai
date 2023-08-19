package echoopenai

import (
	"context"
	"errors"
	"net/http"
)

var ErrUnexpectedModerationRequest = errors.New("error: unexpected input field in moderation request")

type ModerationRequest struct {
	Input any         `json:"input"`
	Model OpenAIModel `json:"model"`
}

type ModerationResponse struct {
	ID      string           `json:"ID"`
	Model   string           `json:"model"`
	Results []map[string]any `json:"results"`
}

func (c *Client) CreateModeration(request ModerationRequest) (response ModerationResponse, err error) {
	return c.CreateModerationWithContext(context.Background(), request)
}

func (c *Client) CreateModerationWithContext(ctx context.Context, request ModerationRequest) (response ModerationResponse, err error) {
	urlSuffix := "moderations"

	if !isString(request.Input) && !isSlice(request.Input) {
		err = ErrUnexpectedModerationRequest
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
