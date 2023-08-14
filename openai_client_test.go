package echoopenai

import (
	"fmt"
	"testing"
)

func TestListModels(t *testing.T) {
	client := NewClient("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")
	models, err := client.ListModels()
	if err != nil {
		t.Error("test ListModes func failed")
	}
	fmt.Printf("%v", models)
}
