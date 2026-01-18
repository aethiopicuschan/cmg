package git

import (
	"bytes"
	"os/exec"

	"iter"
)

// DiffFile represents a single file-level change.
type DiffFile struct {
	Path      string
	Added     int
	Deleted   int
	Diff      string // May be truncated
	Truncated bool
}

// DiffFiles returns a streaming sequence of DiffFile using yield.
//
// The sequence stops automatically when:
//   - the caller stops iteration
//   - the total diff size exceeds MaxTotalBytes
func DiffFiles(opts DiffOptions) iter.Seq[DiffFile] {
	return func(yield func(DiffFile) bool) {
		totalBytes := 0

		for f := range diffNumstatSeq() {
			df := DiffFile{
				Path:    f.Path,
				Added:   f.Added,
				Deleted: f.Deleted,
			}

			if opts.IncludeDiffBody {
				diff, truncated := gitDiffFile(f.Path, opts.MaxPerFileBytes)
				df.Diff = diff
				df.Truncated = truncated

				totalBytes += len(diff)
				if totalBytes > opts.MaxTotalBytes {
					yield(df)
					return
				}
			}

			if !yield(df) {
				return
			}
		}
	}
}

// gitDiffFile returns the diff for a single file, truncated if it exceeds maxBytes.
func gitDiffFile(path string, maxBytes int) (diff string, truncated bool) {
	cmd := exec.Command("git", "diff", "--", path)
	out, err := cmd.Output()
	if err != nil {
		return "", false
	}

	if len(out) <= maxBytes {
		return string(out), false
	}

	var buf bytes.Buffer
	buf.Write(out[:maxBytes])
	buf.WriteString("\n--- DIFF TRUNCATED ---\n")

	return buf.String(), true
}
