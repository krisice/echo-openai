package echoopenai

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestCreateURLFormatImageGeneration(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	req := ImageGenerationRequest{
		Prompt:         "black cat",
		ResponseFormat: ImageFormatURL,
		N:              1,
	}
	res, err := client.CreateImageGeneration(req)
	if err != nil {
		t.Errorf("test generate url format image failed %v", err)
		return
	}
	t.Logf("test generate url format image successed %v", res)
}

func TestCreateB64JSONFormatImageGeneration(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	req := ImageGenerationRequest{
		Prompt:         "color cat",
		ResponseFormat: ImageFormatBase64JSON,
		N:              1,
	}

	res, err := client.CreateImageGeneration(req)
	if err != nil {
		t.Errorf("test generate base64 json format image failed %v", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(res.Data[0].Base64JSON)
	if err != nil {
		t.Errorf("test generate base64 json format image failed %v", err)
		return
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return
	}

	file, err := os.Create("images/image_generation.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return
	}

	t.Logf("test generate base64 json format image successed")
}

func TestImageEdit(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("images/image_generation.png")
	if err != nil {
		t.Errorf("read images/image_generation.png file failed %v", err)
	}

	req := ImageEditRequest{
		Image:  file,
		Prompt: "White dog",
		Config: ImageRequestCommonConfig{
			ResponseFormat: ImageFormatURL,
			N:              1,
			Size:           ImageSize1024x1024,
		},
	}
	res, err := client.CreateImageEdit(req)
	if err != nil {
		t.Errorf("test edit url format image failed %v", err)
		return
	}
	t.Logf("test edit url format image successed %v", res)
}

func TestImageVariation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("images/image_generation.png")
	if err != nil {
		t.Errorf("read images/image_generation.png file failed %v", err)
	}

	req := ImageVariationRequest{
		Image: file,
		Config: ImageRequestCommonConfig{
			ResponseFormat: ImageFormatURL,
			N:              1,
			Size:           ImageSize1024x1024,
		},
	}

	res, err := client.CreateImageVariation(req)
	if err != nil {
		t.Errorf("test variation url format image failed %v", err)
		return
	}
	t.Logf("test variation url format image successed %v", res)
}
