package llm

import (
	"context"

	"github.com/aethiopicuschan/cmg/pkg/config"
	"github.com/cloudwego/eino/components/model"
)

// NewChatModel creates a ToolCallingChatModel from LLMConfig using eino.
//
// It validates the provider and model against SupportedProviders and
// constructs a provider-specific ToolCallingChatModel.
func NewChatModel(cfg config.LLMConfig) (m model.ToolCallingChatModel, err error) {
	if cfg.Models.Chat == nil {
		err = ErrChatModelConfIsMissing
		return
	}

	// Resolve provider definition
	provider, err := findProvider(cfg.Provider)
	if err != nil {
		return
	}

	// Validate model name
	if !contains(provider.ChatModels, cfg.Models.Chat.Model) {
		err = ErrModelIsNotSupported
		return
	}

	// Validate API key if required
	if provider.NeedsAPIKey {
		if cfg.Options["api_key"] == "" {
			err = ErrAPIKeyIsMissing
			return
		}
	}

	switch provider.Name {
	case "openai":
		return newOpenAIChatModel(context.Background(), &cfg)
	default:
		err = ErrProviderNotSupported
		return
	}
}
