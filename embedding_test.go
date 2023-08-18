package echoopenai

import (
	"os"
	"testing"
)

func TestCreateEmbedding(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	req := EmbeddingRequest{
		Model: TextEmbeddingAda002,
		Input: []string{"Test", "Test2"},
	}

	res, err := client.CreateEmbeddings(req)
	if err != nil {
		t.Errorf("test create embedding func failed %v", err)
		return
	}

	t.Logf("test create embedding func successed %v", res)
}
