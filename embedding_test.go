package echoopenai

import "testing"

func TestCreateEmbedding(t *testing.T) {
	client := NewClient("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")
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
