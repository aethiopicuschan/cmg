package git

// DiffOptions controls how git diff is collected for LLM-safe usage.
type DiffOptions struct {
	MaxTotalBytes   int  // Maximum total diff size across all files
	MaxPerFileBytes int  // Maximum diff size per file
	IncludeDiffBody bool // Whether to include actual diff content
}
