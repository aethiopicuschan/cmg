package git

import (
	"bytes"
	"os/exec"
	"strconv"

	"iter"
)

// numstatFile represents a single line from `git diff --numstat`.
type numstatFile struct {
	Path    string
	Added   int
	Deleted int
}

// diffNumstatSeq streams `git diff --numstat` line by line using yield.
func diffNumstatSeq() iter.Seq[numstatFile] {
	return func(yield func(numstatFile) bool) {
		cmd := exec.Command("git", "diff", "--numstat")
		out, err := cmd.Output()
		if err != nil {
			return
		}

		for line := range bytes.SplitSeq(bytes.TrimSpace(out), []byte{'\n'}) {
			fields := bytes.Fields(line)
			if len(fields) < 3 {
				continue
			}

			f := numstatFile{
				Path:    string(fields[2]),
				Added:   atoiSafe(string(fields[0])),
				Deleted: atoiSafe(string(fields[1])),
			}

			if !yield(f) {
				return
			}
		}
	}
}

// atoiSafe converts a string to int and returns 0 on failure.
func atoiSafe(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
