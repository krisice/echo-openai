package echoopenai

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateJSONFormatAuidoTranslation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("audio/audio_translation.m4a")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := AudioTranslationRequest{
		File:           file,
		Model:          Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: AudioJSONFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		t.Errorf("test create auido translation failed %v, response format json", err)
		return
	}

	t.Logf("test create auido translation successed %v, response format json", res)
}

func TestCreateTextFormatAuidoTranslation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("audio/audio_translation.m4a")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := AudioTranslationRequest{
		File:           file,
		Model:          Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: AudioTextFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		t.Errorf("test create auido translation failed %v, response format text", err)
		return
	}

	t.Logf("test create auido translation successed %v, response format text", res)
}

func TestCreateVerboseJSONFormatAuidoTranslation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("audio/audio_translation.m4a")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := AudioTranslationRequest{
		File:           file,
		Model:          Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: AudioVerboseJSONFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		t.Errorf("test create auido translation failed %v, response format verbose json", err)
		return
	}

	t.Logf("test create auido translation successed %v, response format verbose json", res)
}

func TestCreateSRTFormatAuidoTranslation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("audio/audio_translation.m4a")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := AudioTranslationRequest{
		File:           file,
		Model:          Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: AudioSRTFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		t.Errorf("test create auido translation failed %v, response format srt", err)
		return
	}

	t.Logf("test create auido translation successed %v, response format srt", res)
}

func TestCreateVTTFormatAuidoTranslation(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)

	file, err := os.Open("audio/audio_translation.m4a")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := AudioTranslationRequest{
		File:           file,
		Model:          Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: AudioVTTFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		t.Errorf("test create auido translation failed %v, response format vtt", err)
		return
	}

	t.Logf("test create auido translation successed %v, response format vtt", res)
}
