package echoopenai

import (
	"os"
	"testing"
)

func TestCreateModeration(t *testing.T) {
	apiKey := os.Getenv("ECHOOPENAIAPIKEY")
	client := NewClient(apiKey)
	req := ModerationRequest{
		Input: []string{"sex", "xjp"},
		Model: TextModerationLatest,
	}

	res, err := client.CreateModeration(req)
	if err != nil {
		t.Errorf("test create moderation func failed %v", err)
		return
	}

	if len(res.Results) == 0 {
		t.Errorf("test create moderation func failed %v", err)
		return
	}
	t.Logf("test create moderation func successed %v", res)
}
