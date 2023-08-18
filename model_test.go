package echoopenai

import "testing"

func TestListModels(t *testing.T) {
	client := NewClient("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")
	models, err := client.ListModels()
	if err != nil {
		t.Error("test ListModels func failed")
	}
	t.Logf("test ListModels func successed %v", models)
}

func TestRetrieveModel(t *testing.T) {
	client := NewClient("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")
	model, err := client.RetrieveModel(GPT3Dot5Turbo0301)
	if err != nil {
		t.Error("test RetrieveModel func failed")
	}
	t.Logf("test RetrieveModel func successed %v", model)
}
