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
func diffNumstatSeq(opts DiffOptions) iter.Seq[numstatFile] {
	return func(yield func(numstatFile) bool) {
		for _, base := range gitDiffBaseArgs(opts.IgnoreUnstaged) {
			args := append(base, "--numstat")
			cmd := exec.Command("git", args...)
			out, err := cmd.Output()
			if err != nil {
				continue
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
}

// atoiSafe converts a string to int and returns 0 on failure.
func atoiSafe(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func gitDiffBaseArgs(ignoreUnstaged bool) (args [][]string) {
	args = append(args, []string{"diff", "--cached"})
	if !ignoreUnstaged {
		args = append(args, []string{"diff"})
	}
	return
}
