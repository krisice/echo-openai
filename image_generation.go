package echoopenai

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"
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

type ImageRequestCommonConfig struct {
	N              int         `json:"n"`
	Size           ImageSize   `json:"size"`
	ResponseFormat ImageFormat `json:"response_format,omitempty"`
	User           string      `json:"user,omitempty"`
}

type ImageEditRequest struct {
	Image  *os.File                 `json:"image,omitempty"`
	Mask   *os.File                 `json:"mask"`
	Prompt string                   `json:"prompt"`
	Config ImageRequestCommonConfig `json:"config"`
}

type ImageVariationRequest struct {
	Image  *os.File                 `json:"image"`
	Config ImageRequestCommonConfig `json:"config"`
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
	urlSuffix := "images/generations"
	if !checkRequest(request) {
		err = ErrUnexpectedImageRequestType
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

func (c *Client) CreateImageEdit(request ImageEditRequest) (response ImageGenerationResponse, err error) {
	return c.CreateImageEditWithContext(context.Background(), request)
}

func (c *Client) CreateImageEditWithContext(ctx context.Context, request ImageEditRequest) (response ImageGenerationResponse, err error) {
	setDefaultRequestConfig(&request.Config)

	body := &bytes.Buffer{}
	builder := c.createFormBuilder(body)

	err = builder.CreateFormFile("image", request.Image)
	if err != nil {
		return
	}

	if request.Mask != nil {
		err = builder.CreateFormFile("mask", request.Mask)
		if err != nil {
			return
		}
	}

	err = builder.WriteField("prompt", request.Prompt)
	if err != nil {
		return
	}

	err = writeFiled(builder, request.Config)
	if err != nil {
		return
	}

	err = builder.Close()
	if err != nil {
		return
	}

	urlSuffix := "images/edits"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", builder.FormDataContentType())
	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) CreateImageVariation(request ImageVariationRequest) (response ImageGenerationResponse, err error) {
	return c.CreateImageVariationWithContext(context.Background(), request)
}

func (c *Client) CreateImageVariationWithContext(ctx context.Context, request ImageVariationRequest) (response ImageGenerationResponse, err error) {
	setDefaultRequestConfig(&request.Config)

	body := &bytes.Buffer{}
	builder := c.createFormBuilder(body)

	err = builder.CreateFormFile("image", request.Image)
	if err != nil {
		return
	}

	err = writeFiled(builder, request.Config)
	if err != nil {
		return
	}

	err = builder.Close()
	if err != nil {
		return
	}

	urlSuffix := "images/variations"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", builder.FormDataContentType())
	c.setCommonHeader(req)
	err = c.sendRequestWithContext(ctx, req, &response)
	return
}

func setDefaultRequestConfig(request *ImageRequestCommonConfig) {
	if request.N == 0 {
		request.N = 1
	}

	if len(request.Size) == 0 {
		request.Size = ImageSize1024x1024
	}

	if len(request.ResponseFormat) == 0 {
		request.ResponseFormat = ImageFormatURL
	}
}

func writeFiled(builder FormBuilder, config ImageRequestCommonConfig) (err error) {
	err = builder.WriteField("n", strconv.Itoa(config.N))
	if err != nil {
		return
	}

	err = builder.WriteField("size", string(config.Size))
	if err != nil {
		return
	}

	err = builder.WriteField("response_format", string(config.ResponseFormat))
	if err != nil {
		return
	}

	err = builder.WriteField("user", string(config.User))
	if err != nil {
		return
	}

	return
}
