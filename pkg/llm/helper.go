package llm

import (
	"fmt"

	"github.com/aethiopicuschan/cmg/pkg/config"
)

func findProvider(name string) (*config.Provider, error) {
	for _, p := range config.SupportedProviders {
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("unsupported LLM provider: %s", name)
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
