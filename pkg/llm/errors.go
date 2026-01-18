package llm

import "errors"

var (
	ErrGitNotAvailable        = errors.New("git command is not available")
	ErrNotGitRepository       = errors.New("current directory is not a git repository")
	ErrNoGitChanges           = errors.New("no changes detected in the git repository")
	ErrChatModelConfIsMissing = errors.New("chat model configuration is missing")
	ErrModelIsNotSupported    = errors.New("the specified model is not supported")
	ErrAPIKeyIsMissing        = errors.New("api key is missing")
	ErrProviderNotSupported   = errors.New("the specified provider is not supported")
)
