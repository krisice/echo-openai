package echoopenai

import (
	"context"
	"errors"
	"net/http"
)

type ImageSize string
type ImageFormat string

var ErrUnexpectedImageRequestType = errors.New("error: unexpected image request type")

const (
	ImageSize256x256   ImageSize = "256x256"
	ImageSize512x512   ImageSize = "512x512"
	ImageSize1024x1024 ImageSize = "1024x1024"
)

const (
	ImageFormatURL        ImageFormat = "url"
	ImageFormatBase64JSON ImageFormat = "b64_json"
)

type ImageGenerationRequest struct {
	Prompt         string      `json:"prompt"`
	N              int32       `json:"n,omitempty"`
	Size           ImageSize   `json:"size,omitempty"`
	ResponseFormat ImageFormat `json:"response_format,omitempty"`
	User           string      `json:"user,omitempty"`
}
type ImageGenerationResponse struct {
	Created int64             `json:"created"`
	Data    []ImageGeneration `json:"data"`
}
type ImageGeneration struct {
	URL        string `json:"url,omitempty"`
	Base64JSON string `json:"b64_json,omitempty"`
}
func (c *Client) CreateImageGeneration(request ImageGenerationRequest) (response ImageGenerationResponse, err error) {
	return c.CreateImageGenerationWithContext(context.Background(), request)
}

func (c *Client) CreateImageGenerationWithContext(ctx context.Context, request ImageGenerationRequest) (response ImageGenerationResponse, err error) {
	if request.N == 0 {
		request.N = 1
	}

	urlSuffix := "images/generations"
	c.generateImageWithContext(ctx, urlSuffix, request, &response)
	return
}
func (c *Client) generateImageWithContext(ctx context.Context, urlSuffix string, request, response any) (err error) {
	if !checkRequest(request) {
		err = ErrUnexpectedImageRequestType
		return
	}

	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), request)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	c.sendRequestWithContext(ctx, req, response)
	return
}
