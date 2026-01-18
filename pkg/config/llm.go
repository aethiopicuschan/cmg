package config

type Provider struct {
	Name        string   // e.g., openai, gemini, claude
	Description string   // e.g., OpenAI, Google Gemini, Anthropic Claude
	ChatModels  []string // e.g., gpt-4o, gpt-4o-mini
	IsLocal     bool     // e.g., ollama
	NeedsAPIKey bool     // whether this provider needs an API key
}

var SupportedProviders = []Provider{
	{
		Name:        "openai",
		Description: "OpenAI",
		ChatModels: []string{
			"gpt-4o",
			"gpt-4o-mini",
			"gpt-4.1",
			"gpt-3.5-turbo",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "gemini",
		Description: "Google Gemini",
		ChatModels: []string{
			"gemini-1.5-pro",
			"gemini-1.5-flash",
			"gemini-1.0-pro",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "claude",
		Description: "Anthropic Claude",
		ChatModels: []string{
			"claude-3-opus",
			"claude-3-sonnet",
			"claude-3-haiku",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "openrouter",
		Description: "OpenRouter (multi-provider gateway)",
		ChatModels: []string{
			"openai/gpt-4o",
			"openai/gpt-4o-mini",
			"anthropic/claude-3-sonnet",
			"google/gemini-1.5-pro",
			"meta-llama/llama-3-70b",
			"qwen/qwen2.5-72b-instruct",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "deepseek",
		Description: "DeepSeek",
		ChatModels: []string{
			"deepseek-chat",
			"deepseek-coder",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "qwen",
		Description: "Alibaba Qwen",
		ChatModels: []string{
			"qwen-turbo",
			"qwen-plus",
			"qwen-max",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "qianfan",
		Description: "Baidu Qianfan (ERNIE)",
		ChatModels: []string{
			"ernie-bot",
			"ernie-bot-turbo",
			"ernie-bot-4",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "ark",
		Description: "ByteDance Ark",
		ChatModels: []string{
			"doubao-lite",
			"doubao-pro",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "arkbot",
		Description: "ByteDance Ark (Bot API)",
		ChatModels: []string{
			"doubao-lite",
			"doubao-pro",
		},
		NeedsAPIKey: true,
	},

	{
		Name:        "ollama",
		Description: "Ollama (local LLM)",
		ChatModels: []string{
			"llama3",
			"llama3:8b",
			"llama3:70b",
			"mistral",
			"mixtral",
			"qwen2.5",
			"qwen2.5:14b",
		},
		IsLocal: true,
	},
}

func ProviderDescriptions() (descs []string) {
	for _, p := range SupportedProviders {
		descs = append(descs, p.Description)
	}
	return
}
