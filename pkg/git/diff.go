package git

// DiffOptions controls how git diffs are collected and formatted
// for LLM-safe commit message generation.
type DiffOptions struct {
	IgnoreUnstaged  bool // Whether to ignore unstaged (working tree) changes
	MaxTotalBytes   int  // Maximum total diff size across all files
	MaxPerFileBytes int  // Maximum diff size per file
	OutputDetails   bool // Whether to generate a multi-line (detailed) commit message
	IncludeDiffBody bool // Whether to include full diff hunks instead of summaries only
}
