package llm

import (
	"context"

	"github.com/aethiopicuschan/cmg/pkg/config"
	"github.com/cloudwego/eino-ext/components/model/openai"
)

func newOpenAIChatModel(ctx context.Context, cfg *config.LLMConfig) (m *openai.ChatModel, err error) {
	key := cfg.Options["api_key"].(string)
	c := openai.ChatModelConfig{
		APIKey: key,
		Model:  cfg.Models.Chat.Model,
	}
	m, err = openai.NewChatModel(ctx, &c)
	return
}
