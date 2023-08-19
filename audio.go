package echoopenai

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
)

type AudioFormat string

const (
	AudioJSONFormat        AudioFormat = "json"
	AudioTextFormat        AudioFormat = "text"
	AudioSRTFormat         AudioFormat = "srt"
	AudioVerboseJSONFormat AudioFormat = "verbose_json"
	AudioVTTFormat         AudioFormat = "vtt"
)

type AudioTranslationRequest struct {
	File           *os.File    `json:"file"`
	Model          OpenAIModel `json:"model"`
	Prompt         string      `json:"prompt"`
	ResponseFormat AudioFormat `json:"response_format"`
	Temperature    float32     `json:"temperature"`
	Language       string      `json:"language"`
}

type AudiTranslationResponse struct {
	Text string
}

func (c *Client) CreateAudioTranscription(request AudioTranslationRequest) (response AudiTranslationResponse, err error) {
	return c.CreateAudioTranscriptionWithContext(context.Background(), request)
}

func (c *Client) CreateAudioTranscriptionWithContext(ctx context.Context, request AudioTranslationRequest) (response AudiTranslationResponse, err error) {
	urlSuffix := "audio/translations"

	body := &bytes.Buffer{}
	builder := c.createFormBuilder(body)

	err = builder.CreateFormFile("file", request.File)
	if err != nil {
		return
	}

	err = builder.WriteField("model", string(request.Model))
	if err != nil {
		return
	}

	err = builder.WriteField("prompt", request.Prompt)
	if err != nil {
		return
	}

	err = builder.WriteField("response_format", string(request.ResponseFormat))
	if err != nil {
		return
	}

	err = builder.WriteField("temperature", fmt.Sprintf("%.2f", request.Temperature))
	if err != nil {
		return
	}

	if len(request.Language) != 0 {
		urlSuffix = "audio/transcriptions"
		err = builder.WriteField("language", request.Language)
		if err != nil {
			return
		}
	}

	err = builder.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.getFullURL(urlSuffix), body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", builder.FormDataContentType())
	c.setCommonHeader(req)

	if request.HasJSONResponse() {
		err = c.sendRequestWithContext(ctx, req, &response)
	} else {
		err = c.sendRequestWithContext(ctx, req, &response.Text)
	}
	return
}

func (r AudioTranslationRequest) HasJSONResponse() bool {
	return r.ResponseFormat == "" || r.ResponseFormat == AudioJSONFormat
}
