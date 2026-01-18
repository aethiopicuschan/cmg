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

func HasDiff(opts DiffOptions) bool {
	for _, base := range gitDiffBaseArgs(opts.IgnoreUnstaged) {
		args := append(base, "--quiet")
		cmd := exec.Command("git", args...)
		err := cmd.Run()

		if err == nil {
			continue
		}
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return true
		}
	}
	return false
}
