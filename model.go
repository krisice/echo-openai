package echoopenai

import (
	"context"
	"fmt"
	"net/http"
)

type OpenAIModel string

const (
	Ada                       OpenAIModel = "ada"
	AdaCodeSearchCode         OpenAIModel = "ada-code-search-code"
	AdaSearchDocument         OpenAIModel = "ada-search-document"
	AdaCodeSearchText         OpenAIModel = "ada-code-search-text"
	AdaSearchQuery            OpenAIModel = "ada-search-query"
	AdaSimilarity             OpenAIModel = "ada-similarity"
	Babbage                   OpenAIModel = "babbage"
	BabbageCodeSearchCode     OpenAIModel = "babbage-code-search-code"
	BabbageCodeSearchText     OpenAIModel = "babbage-code-search-text"
	BabbageSearchDocument     OpenAIModel = "babbage-search-document"
	BabbageSearchQuery        OpenAIModel = "babbage-search-query"
	BabbageSimilarity         OpenAIModel = "babbage-similarity"
	CodeDavinciEdit001        OpenAIModel = "code-davinci-edit-001"
	CodeSearchAdaCode001      OpenAIModel = "code-search-ada-code-001"
	CodeSearchAdaText001      OpenAIModel = "code-search-ada-text-001"
	CodeSearchBabbageCode001  OpenAIModel = "code-search-babbage-code-001"
	TextSearchBabbageDoc001   OpenAIModel = "text-search-babbage-doc-001"
	CodeSearchBabbageText001  OpenAIModel = "code-search-babbage-text-001"
	Curie                     OpenAIModel = "curie"
	CurieInstructBeta         OpenAIModel = "curie-instruct-beta"
	CurieSearchDocument       OpenAIModel = "curie-search-document"
	CurieSearchQuery          OpenAIModel = "curie-search-query"
	CurieSimilarity           OpenAIModel = "curie-similarity"
	Davinci                   OpenAIModel = "davinci"
	DavinciInstructBeta       OpenAIModel = "davinci-instruct-beta"
	DavinciSearchDocument     OpenAIModel = "davinci-search-document"
	DavinciSearchQuery        OpenAIModel = "davinci-search-query"
	DavinciSimilarity         OpenAIModel = "davinci-similarity"
	GPT3Dot5Turbo             OpenAIModel = "gpt-3.5-turbo"
	GPT3Dot5Turbo0301         OpenAIModel = "gpt-3.5-turbo-0301"
	GPT3Dot5Turbo0613         OpenAIModel = "gpt-3.5-turbo-0613"
	GPT3Dot5Turbo16k          OpenAIModel = "gpt-3.5-turbo-16k"
	GPT3Dot5Turbo16K0613      OpenAIModel = "gpt-3.5-turbo-16k-0613"
	TextAda001                OpenAIModel = "text-ada-001"
	TextBabbage001            OpenAIModel = "text-babbage-001"
	TextCurie001              OpenAIModel = "text-curie-001"
	TextDavinci002            OpenAIModel = "text-davinci-002"
	TextDavinci001            OpenAIModel = "text-davinci-001"
	TextDavinci003            OpenAIModel = "text-davinci-003"
	TextDavinciEdit001        OpenAIModel = "text-davinci-edit-001"
	TextEmbeddingAda002       OpenAIModel = "text-embedding-ada-002"
	TextModerationLatest      OpenAIModel = "text-moderation-latest"
	TextModerationStable      OpenAIModel = "text-moderation-stable"
	TextSimilarityAda001      OpenAIModel = "text-similarity-ada-001"
	TextSimilarityBabbage001  OpenAIModel = "text-similarity-babbage-001"
	TextSimilarityCurie001    OpenAIModel = "text-similarity-curie-001"
	TextSimilarityDavinci001  OpenAIModel = "text-similarity-davinci-001"
	TextSearchAdaDoc001       OpenAIModel = "text-search-ada-doc-001"
	TextSearchAdaQuery001     OpenAIModel = "text-search-ada-query-001"
	TextSearchBabbageQuery001 OpenAIModel = "text-search-babbage-query-001"
	TextSearchCurieDoc001     OpenAIModel = "text-search-curie-doc-001"
	TextSearchCurieQuery001   OpenAIModel = "text-search-curie-query-001"
	TextSearchDavinciDoc001   OpenAIModel = "text-search-davinci-doc-001"
	TextSearchDavinciQuery001 OpenAIModel = "text-search-davinci-query-001"
	Whisper1                  OpenAIModel = "whisper-1"
)

type ModelPermission struct {
	ID                 string `json:"id"`
	Object             string `json:"object"`
	Created            int64  `json:"created"`
	AllowCreateEngine  bool   `json:"allow create_engine"`
	AllowSampling      bool   `json:"allow_sampling"`
	AllowLogprobs      bool   `json:"allow_logprobs"`
	AllowSearchIndices bool   `json:"allow_search_indices"`
	AllowView          bool   `json:"allow_view"`
	AllowFineTuning    bool   `json:"allow_fine_tuning"`
	Organization       string `json:"organization"`
	Group              string `json:"group"`
	IsBlocking         bool   `json:"is_blocking"`
}

type ModelEntry struct {
	ID         string            `json:"id"`
	Object     string            `json:"object"`
	Created    int64             `json:"created"`
	OwnedBy    string            `json:"owned_by"`
	Permission []ModelPermission `json:"permission"`
	Root       string            `json:"root"`
	Parent     string            `json:"parent"`
}

type ListModelsResponse struct {
	Data   []ModelEntry `json:"data"`
	Object string       `json:"object"`
}

func (c *Client) ListModels() (ListModelsResponse, error) {
	return c.ListModelsWithContext(context.Background())
}

func (c *Client) ListModelsWithContext(ctx context.Context) (response ListModelsResponse, err error) {
	urlSuffix := "models"
	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodGet, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	c.sendRequestWithContext(ctx, req, &response)
	return
}

func (c *Client) RetrieveModel(id OpenAIModel) (ModelEntry, error) {
	return c.RetrieveModelWithContext(context.Background(), id)
}

func (c *Client) RetrieveModelWithContext(ctx context.Context, id OpenAIModel) (response ModelEntry, err error) {
	urlSuffix := fmt.Sprintf("%v/%v", "models", id)
	req, err := c.requestBuilder.BuildWithContext(ctx, http.MethodGet, c.getFullURL(urlSuffix), nil)
	if err != nil {
		return
	}

	c.setCommonHeader(req)
	c.sendRequestWithContext(ctx, req, &response)
	return
}
