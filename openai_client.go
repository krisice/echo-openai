package echoopenai

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	apiKey   string
	Organize string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (c *Client) ListModels() (*ListModelsResponse, error) {
	res, err := http.Get("https://api.openai.com/v1/models")
	if err != nil {
		return nil, err
	}
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+ c.apiKey)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var models ListModelsResponse
	err = json.Unmarshal(body, &models)
	if err != nil {
		return nil, err
	}
	return &models, nil
}
