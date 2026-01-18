package llm

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aethiopicuschan/cmg/pkg/git"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// GenerateCommitMessage uses an LLM (via eino) to generate a commit message
// based on the current git diff.
func GenerateCommitMessage(ctx context.Context, llm model.ToolCallingChatModel, opts git.DiffOptions) (string, error) {
	if !git.IsGitAvailable() {
		return "", ErrGitNotAvailable
	}
	if !git.IsGitRepository() {
		return "", ErrNotGitRepository
	}
	if !git.HasDiff(opts) {
		return "", ErrNoGitChanges
	}

	messages := buildCommitPrompt(opts, opts.OutputDetails)

	// No tools are bound for this request.
	resp, err := llm.Generate(ctx, messages)
	if err != nil {
		return "", err
	}

	return resp.Content, nil
}

func buildCommitPrompt(
	opts git.DiffOptions,
	details bool,
) []*schema.Message {
	var msgs []*schema.Message

	msgs = append(msgs, schema.SystemMessage(
		"You are an expert software engineer. "+
			"Generate a concise and informative git commit message "+
			"based on the given changes.",
	))

	// File-level summary
	var summary bytes.Buffer
	summary.WriteString("File-level change summary:\n")

	for f := range git.DiffFiles(opts) {
		summary.WriteString(
			fmt.Sprintf(
				"- %s (+%d -%d)\n",
				f.Path,
				f.Added,
				f.Deleted,
			),
		)
	}
	msgs = append(msgs, schema.UserMessage(summary.String()))

	// Hunk-level details (only if enabled)
	if opts.IncludeDiffBody {
		var hunks bytes.Buffer
		hunks.WriteString("Detailed changes (diff hunks):\n")

		for h := range git.DiffHunks(opts) {
			hunks.WriteString(
				fmt.Sprintf(
					"\nFile: %s\n%s\n%s",
					h.FilePath,
					h.Header,
					h.Body,
				),
			)
		}
		msgs = append(msgs, schema.UserMessage(hunks.String()))
	}

	// Output format instruction (IMPORTANT)
	if details {
		msgs = append(msgs, schema.UserMessage(
			"Write a git commit message with the following format:\n\n"+
				"- First line: a concise commit title (imperative mood)\n"+
				"- Second line: empty\n"+
				"- Following lines: a short commit body explaining WHAT and WHY\n\n"+
				"Keep the message clear and professional. "+
				"If applicable, follow Conventional Commits.",
		))
	} else {
		msgs = append(msgs, schema.UserMessage(
			"Write a single-line git commit message (title only). "+
				"Use the imperative mood and keep it concise. "+
				"Do not include any additional lines.",
		))
	}

	return msgs
}
