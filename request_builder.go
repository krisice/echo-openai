package echoopenai

import (
	"bytes"
	"context"
	"net/http"
)

type RequestBuild interface {
	Build(ctx context.Context, method, url string, request any) (*http.Request, error)
}

type HTTPRequestBuilder struct {
	marshaller Marshaller
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		marshaller: &JSONMarshaller{},
	}
}

func (b *HTTPRequestBuilder) build(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	reqBytes, err := b.marshaller.Marshal(request)
	if err != nil {
		return nil, err
	}

	return http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBytes),
	)
}
