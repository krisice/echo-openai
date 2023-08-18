package echoopenai

import (
	"os"
	"testing"
)

func TestListModels(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	models, err := client.ListModels()
	if err != nil {
		t.Error("test ListModels func failed")
	}
	t.Logf("test ListModels func successed %v", models)
}

func TestRetrieveModel(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	model, err := client.RetrieveModel(GPT3Dot5Turbo0301)
	if err != nil {
		t.Error("test RetrieveModel func failed")
	}
	t.Logf("test RetrieveModel func successed %v", model)
}
