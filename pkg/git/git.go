package git

import (
	"os"
	"os/exec"
	"path/filepath"
)

// IsGitAvailable checks whether the git command is available in PATH.
func IsGitAvailable() (available bool) {
	_, err := exec.LookPath("git")
	return err == nil
}

// IsGitRepository checks if the current directory is inside a Git repository.
//
// It works by walking up from the current working directory to the filesystem
// root and checking for the presence of a ".git" directory or file.
// This approach supports normal repositories, submodules, and worktrees
// without relying on the git command.
func IsGitRepository() (is bool) {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return false
	}

	for {
		// Construct the path to ".git"
		gitPath := filepath.Join(dir, ".git")

		info, err := os.Stat(gitPath)
		if err == nil {
			// If ".git" exists, this directory is inside a Git repository.
			// ".git" can be either:
			// - a directory (standard repository)
			// - a file (worktree or submodule)
			if info.IsDir() || info.Mode().IsRegular() {
				return true
			}
		}

		// Move up to the parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached the filesystem root without finding ".git"
			return false
		}
		dir = parent
	}
}

// HasDiff checks whether there are any uncommitted changes.
//
// It returns true if there is at least one diff, false otherwise.
// Internally, it uses `git diff --quiet`, which is efficient and
// does not generate diff output.
func HasDiff() bool {
	cmd := exec.Command("git", "diff", "--quiet")
	err := cmd.Run()

	// Exit status:
	//   0 => no diff
	//   1 => diff exists
	//   other => error
	if err == nil {
		return false
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode() == 1
	}

	// Any other error (e.g. git not available)
	return false
}
