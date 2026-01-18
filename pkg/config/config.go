package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aethiopicuschan/cmg/pkg/internal/colors"
	"github.com/aethiopicuschan/cmg/pkg/internal/prompt"
	"github.com/aethiopicuschan/cmg/pkg/value/constants"
	"github.com/aethiopicuschan/config-go"
)

// Singleton config instance
var conf *Config

// Returns true if the config file exists
func IsConfigExists() (exist bool, err error) {
	exist, err = config.DirExists(constants.AppName)
	return
}

// Returns the config directory path
func GetConfigDir() (dir string, err error) {
	dir, err = config.GetConfigDir(constants.AppName)
	return
}

// Returns true if the config directory exists
func IsConfigDirExists() (exist bool, err error) {
	exist, err = config.DirExists(constants.AppName)
	return
}

// GetConfig returns the loaded configuration
func GetConfig() (c *Config, err error) {
	if conf == nil {
		if err = LoadConfig(); err != nil {
			return
		}
	}
	c = conf
	return
}

// Loads configuration from file (creates one if missing)
func LoadConfig() (err error) {
	dir, err := config.GetConfigDir(constants.AppName)
	if err != nil {
		return
	}

	configPath := filepath.Join(dir, constants.ConfigFileName)

	// Create config file if it does not exist
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		if err = EnsureConfig(); err != nil {
			return
		}
	}

	cj := config.NewConfig(configPath)
	if err = cj.Load(); err != nil {
		return
	}

	buf, err := cj.Read()
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, &conf); err != nil {
		return
	}

	return
}

// Runs the interactive configuration wizard and writes config file
func EnsureConfig() (err error) {
	// Ensure config directory
	_, err = GetConfigDir()
	if err != nil {
		return
	}

	fmt.Println(colors.Green("✨ Initial LLM Configuration"))
	fmt.Println(colors.Blue("This wizard will help you create an initial configuration.\n"))

	// --- Provider selection ---
	providerIndex := prompt.PromptSelect(
		colors.Cyan("Select LLM provider"),
		ProviderDescriptions(),
		0,
	)
	provider := SupportedProviders[providerIndex]

	fmt.Println(
		colors.Yellow("Selected provider: ") +
			colors.Green(provider.Name) +
			colors.Blue(" ("+provider.Description+")"),
	)

	// --- Chat model selection ---
	modelIndex, customModel := prompt.PromptSelectOrInputIndex(
		colors.Cyan("Select chat model"),
		colors.Cyan("Input chat model"),
		provider.ChatModels,
		provider.ChatModels[0],
	)
	var model string
	if modelIndex == -1 {
		model = customModel
	} else {
		model = provider.ChatModels[modelIndex]
	}

	// --- Streaming option ---
	streaming := prompt.PromptBool(
		colors.Cyan("Enable streaming"),
		provider.IsLocal,
	)

	// --- Provider-specific options ---
	options := map[string]any{}

	if provider.IsLocal {
		fmt.Println(colors.Magenta("\nLocal LLM settings"))
		options["endpoint"] = prompt.PromptString(
			colors.Cyan("Endpoint"),
			"http://localhost:11434",
		)
	} else if provider.NeedsAPIKey {
		var ok bool
		fmt.Println(colors.Magenta("API authentication"))
		for !ok {
			options["api_key"] = prompt.PromptSecretString(
				colors.Cyan("API key"),
			)
			if options["api_key"] != "" {
				ok = true
			}
		}
	}

	// --- Build config ---
	conf = &Config{
		LLM: LLMConfig{
			Provider: provider.Name,
			Models: LLMModelConfig{
				Chat: &ChatModelConfig{
					Model:     model,
					Streaming: streaming,
				},
			},
			Options: options,
		},
	}

	// --- Write config file ---
	dir, err := config.EnsureConfigDir(constants.AppName)
	if err != nil {
		return
	}

	configPath := filepath.Join(dir, constants.ConfigFileName)

	b, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return
	}

	if err = os.WriteFile(configPath, b, 0o600); err != nil {
		return
	}

	fmt.Println()
	fmt.Println(colors.Green("✅ Configuration file created successfully"))
	fmt.Println(colors.Blue("Path: ") + colors.Yellow(configPath))

	return
}

// Initializes the base config directory (~/.config)
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %V", err)
	}
	configDirPath := filepath.Join(home, ".config")
	config.SetConfigDir(configDirPath)
}
