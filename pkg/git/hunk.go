package git

import (
	"bytes"
	"os/exec"

	"iter"
)

// DiffHunk represents a single diff hunk (@@ ... @@).
type DiffHunk struct {
	FilePath string
	Header   string // @@ -a,b +c,d @@
	Body     string // Lines belonging to this hunk
}

// DiffHunks returns a streaming sequence of DiffHunk.
//
// Diff is split by hunk (@@) boundaries, allowing the caller to stop
// safely at semantically meaningful points.
func DiffHunks(opts DiffOptions) iter.Seq[DiffHunk] {
	return func(yield func(DiffHunk) bool) {
		totalBytes := 0

		for f := range diffNumstatSeq() {
			for h := range diffHunksFromFile(f.Path, opts.MaxPerFileBytes) {
				size := len(h.Header) + len(h.Body)
				totalBytes += size

				if totalBytes > opts.MaxTotalBytes {
					yield(h)
					return
				}

				if !yield(h) {
					return
				}
			}
		}
	}
}

// diffHunksFromFile splits a file diff into hunks and yields them incrementally.
func diffHunksFromFile(path string, maxBytes int) iter.Seq[DiffHunk] {
	return func(yield func(DiffHunk) bool) {
		cmd := exec.Command("git", "diff", "--unified=0", "--", path)
		out, err := cmd.Output()
		if err != nil {
			return
		}

		var (
			currentHeader string
			body          bytes.Buffer
			usedBytes     int
		)

		for line := range bytes.SplitSeq(out, []byte{'\n'}) {
			if bytes.HasPrefix(line, []byte("@@")) {
				if currentHeader != "" {
					if !yield(DiffHunk{
						FilePath: path,
						Header:   currentHeader,
						Body:     body.String(),
					}) {
						return
					}
				}

				currentHeader = string(line)
				body.Reset()
				continue
			}

			if currentHeader == "" {
				// Skip diff metadata before first hunk
				continue
			}

			if usedBytes+len(line) > maxBytes {
				body.WriteString("\n--- HUNK TRUNCATED ---\n")
				yield(DiffHunk{
					FilePath: path,
					Header:   currentHeader,
					Body:     body.String(),
				})
				return
			}

			body.Write(line)
			body.WriteByte('\n')
			usedBytes += len(line)
		}

		if currentHeader != "" {
			yield(DiffHunk{
				FilePath: path,
				Header:   currentHeader,
				Body:     body.String(),
			})
		}
	}
}
