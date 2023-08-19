package echoopenai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const OpenAIRequestAPIv1 = "https://api.openai.com/v1"
const defaultEmptyMessagesLimit uint = 300

type TokenUsage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type Client struct {
	apiKey            string
	baseURL           string
	client            *http.Client
	requestBuilder    RequestBuilder
	createFormBuilder func(io.Writer) FormBuilder
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:         apiKey,
		baseURL:        OpenAIRequestAPIv1,
		client:         &http.Client{},
		requestBuilder: NewHTTPRequestBuilder(),
		createFormBuilder: func(body io.Writer) FormBuilder {
			return NewFormBuilder(body)
		},
	}
}

func (c *Client) setCommonHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
}

func (c *Client) setStreamHeader(req *http.Request) {
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
}

// func (c *Client) setFormDataHeader(req *http.Request) {
// 	req.Header.Set("Content-Type", "multipart/form-data")
// }

func (c *Client) getFullURL(suffix string) string {
	return fmt.Sprintf("%v/%v", c.baseURL, suffix)
}

func (c *Client) sendRequestWithContext(ctx context.Context, req *http.Request, v any) error {
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}

	return decodeResponse(res.Body, v)
}

func (c *Client) createStreamRequestWithContext(ctx context.Context, method string, urlSuffix string, body any) (*http.Request, error) {
	req, err := c.requestBuilder.BuildWithContext(ctx, method, c.getFullURL(urlSuffix), body)
	if err != nil {
		return nil, err
	}

	c.setStreamHeader(req)
	c.setCommonHeader(req)
	return req, nil
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &RequestError{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	return errRes.Error
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return decodeString(body, result)
	}
	return json.NewDecoder(body).Decode(v)
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func isString(v any) bool {
	_, ok := v.(string)
	return ok
}

func isSlice(v any) bool {
	_, ok := v.([]string)
	return ok
}

func checkRequest(request any) bool {
	switch request.(type) {
	case ImageGenerationRequest, ImageEditRequest, ImageVariationRequest:
		return true
	default:
		return false
	}
}
