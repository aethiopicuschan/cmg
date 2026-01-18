package config

type Config struct {
	LLM LLMConfig `json:"llm"`
}

type LLMConfig struct {
	Provider string         `json:"provider"`
	Models   LLMModelConfig `json:"models"`
	Options  map[string]any `json:"options,omitempty"`
}

type LLMModelConfig struct {
	Chat      *ChatModelConfig      `json:"chat,omitempty"`
	Embedding *EmbeddingModelConfig `json:"embedding,omitempty"`
}

type ChatModelConfig struct {
	Model        string   `json:"model"`
	Temperature  *float32 `json:"temperature,omitempty"`
	MaxTokens    *int     `json:"max_tokens,omitempty"`
	SystemPrompt string   `json:"system_prompt,omitempty"`
	Streaming    bool     `json:"streaming,omitempty"`
}

type EmbeddingModelConfig struct {
	Model      string `json:"model"`
	Dimensions *int   `json:"dimensions,omitempty"`
}
