# echo-openai

This library provides unofficial Go clients for [OpenAI API](https://platform.openai.com/). We support:

- Chat Completion
- Audio Transaction
- Image Generation
- Embedding
- Moderation
- File Upload

## Usage

### Chat Completion example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)
	
	message := openai.ChatCompletionMessage{
		Role:    openai.ChatCompletionRoleUser,
		Content: "Test",
	}
	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, message)
	
	request := openai.ChatCompletionRequest{
		Model:    string(openai.GPT3Dot5Turbo),
		Messages: messages,
	}
	
	res, err := client.CreateChatCompletion(request)
	if err != nil {
		//handle error
		return
	}
	
	if len(res.Choices) != 0 {
		message = res.Choices[0].Message
	}
}

```

### Chat Completion Stream example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)
	
	message := openai.ChatCompletionMessage{
		Role:    openai.ChatCompletionRoleUser,
		Content: "Test",
	}
	messages := make([]ChatCompletionMessage, 0)
	messages = append(messages, message)
	
	request := openai.ChatCompletionRequest{
		Model:    string(openai.GPT3Dot5Turbo),
		Messages: messages,
		Stream:   true,
	}
	
	stream, err := client.CreateChatCompletionStream(request)
	if err != nil {
		fmt.Error("test CreateChatCompletionStream func failed")
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("stream recv completion")
			} else {
				fmt.Println("stream recv failed %v", err.Error())
			}
			break
		}
		
		if len(res.Choices) != 0 {
			message = res.Choices[0].Message
		}
	}
}
```

### Audio Translation example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)

	file, err := os.Open("your audio file path")
	if err != nil {
		fmt.Printf("open audio file failed %v", err)
		return
	}

	req := openai.AudioTranslationRequest{
		File:           file,
		Model:          openai.Whisper1,
		Prompt:         "this is new audio",
		ResponseFormat: openai.AudioJSONFormat,
	}

	res, err := client.CreateAudioTranscription(req)
	if err != nil {
		fmt.Errorf("auido translation failed %v, response format json", err)
		return
	}
	
	fmt.Printf("auido translation successed %v, response format json", res.Text)
}
```

### Image Generation example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)

	req := openai.ImageGenerationRequest{
		Prompt:         "black cat",
		ResponseFormat: openai.ImageFormatURL,
		N:              1,
	}
	
	res, err := client.CreateImageGeneration(req)
	if err != nil {
		fmt.Errorf("generate url format image failed %v", err)
		return
	}
	if len(res.Data) != 0 {
		fmt.Printf("generated image URL is %v", res.Data[0].URL)
	}
}
```

### Embedding example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)
	req := openai.EmbeddingRequest{
		Model: openai.TextEmbeddingAda002,
		Input: []string{"Test", "Test2"},
	}

	res, err := client.CreateEmbeddings(req)
	if err != nil {
		fmt.Errorf("create embedding failed %v", err)
		return
	}
	
	if len(res.Data) != 0 {
		fmt.Printf("embedding response %v", res.Data[0].Embedding)
	}
}
```

### Moderation example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)
	
	req := openai.ModerationRequest{
		Input: []string{"sex", "xjp"},
		Model: openai.TextModerationLatest,
	}

	res, err := client.CreateModeration(req)
	if err != nil {
		fmt.Errorf("create moderation failed %v", err)
		return
	}

	if len(res.Results) != 0 {
		fmt.Errorf("create moderation failed %v", err)
		return
	}
	fmt.Printf("create moderation result is %v", res.Results)
}
```

### File Upload example usage

```
package main

import (
	"context"
	"fmt"
	openai "github.com/kriskice/echo-openai"
)

func main() {
	apiKey := "your openai api key"
	client := openai.NewClient(apiKey)
	
	//open your jsonl format data file
	file, err := os.Open("your data file path")
	if err != nil {
		t.Errorf("open prepared upload files failed %v", err)
		return
	}

	req := FileUploadRequest{
		File:    file,
		Purpose: "fine-tune",
	}

	res, err := client.UploadFile(req)
	if err != nil {
		fmt.Errorf("upload file failed %v", err)
		return
	}
	
	if len(res.Data) != 0 {
		fmt.Printf("upload file status %v", res.Data[0].StatusDetails)
	}
}
```

## TODO

- Provide API for model fine tune
- Provide a method to count token usage
- Unit test completion