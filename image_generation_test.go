package echoopenai

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestCreateImageGeneration(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	req := ImageGenerationRequest{
		Prompt:         "black cat",
		ResponseFormat: ImageFormatURL,
	}
	res, err := client.CreateImageGeneration(req)
	if err != nil {
		t.Errorf("test generate url format image failed %v", err)
	}
	t.Logf("test generate url format image successed")

	req = ImageGenerationRequest{
		Prompt:         "black cat",
		ResponseFormat: ImageFormatBase64JSON,
	}

	res, err = client.CreateImageGeneration(req)
	if err != nil {
		t.Errorf("test generate base64 json format image failed %v", err)
	}

	bytes, err := base64.StdEncoding.DecodeString(res.Data[0].Base64JSON)
	if err != nil {
		t.Errorf("test generate base64 json format image failed %v", err)
	}

	err = os.WriteFile("images/image_generation.png", bytes, 0666)
	if err != nil {
		t.Errorf("test generate base64 json format image failed %v", err)
	}

	t.Logf("test generate base64 json format image successed")
}
