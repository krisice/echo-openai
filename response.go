package echoopenai

type ListModelEntry struct {
	ID         string
	Object     string
	OwnedBy    string
	Permission []string
}

type ListModelsResponse struct {
	Models []ListModelEntry
	Object string
}
