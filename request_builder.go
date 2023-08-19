package echoopenai

import (
	"bytes"
	"context"
	"net/http"
)

type RequestBuilder interface {
	Build(method, url string, request any) (*http.Request, error)
	BuildWithContext(ctx context.Context, method, url string, request any) (*http.Request, error)
}

type HTTPRequestBuilder struct {
	marshaller Marshaller
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		marshaller: &JSONMarshaller{},
	}
}

func (b *HTTPRequestBuilder) BuildWithContext(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	var err error

	var contentType string
	switch request.(type) {
	case ImageEditRequest, ImageVariationRequest:
		reqBytes, contentType, err = generateFormData(request)
	default:
		reqBytes, err = b.marshaller.Marshal(request)
	}

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	)
}

func (b *HTTPRequestBuilder) Build(method, url string, request any) (*http.Request, error) {
	return b.BuildWithContext(context.Background(), method, url, request)
}
