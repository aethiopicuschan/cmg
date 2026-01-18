package cmd

import (
	"fmt"
	"runtime/debug"

	pkgConfig "github.com/aethiopicuschan/cmg/pkg/config"
	"github.com/aethiopicuschan/cmg/pkg/git"
	"github.com/aethiopicuschan/cmg/pkg/llm"
	"github.com/spf13/cobra"
)

var (
	details        bool
	ignoreUnstaged bool
)

var rootCmd = &cobra.Command{
	Use:           "cmg",
	Long:          `Commit message generator based on git diff using an LLM`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Load configuration
		cfg, err := pkgConfig.GetConfig()
		if err != nil {
			return
		}

		// Ensure chat model config exists
		if cfg.LLM.Models.Chat == nil {
			err = fmt.Errorf("chat model configuration is required")
			return
		}

		// Create LLM (ToolCallingChatModel)
		chatModel, err := llm.NewChatModel(cfg.LLM)
		if err != nil {
			return
		}

		// Build diff options (LLM-safe defaults)
		diffOpts := git.DiffOptions{
			MaxTotalBytes:   120_000,
			MaxPerFileBytes: 8_000,
			IncludeDiffBody: true,
			IgnoreUnstaged:  ignoreUnstaged,
		}

		ctx := cmd.Context()

		// Generate commit message
		message, err := llm.GenerateCommitMessage(
			ctx,
			chatModel,
			diffOpts,
		)
		if err != nil {
			return
		}

		// Output result
		fmt.Fprintln(cmd.OutOrStdout(), message)
		return
	},
}

func Execute() (err error) {
	return rootCmd.Execute()
}

func init() {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		rootCmd.Version = bi.Main.Version
	}
	rootCmd.Flags().BoolVarP(
		&details,
		"details",
		"d",
		false,
		"include detailed commit body (multi-line commit message)",
	)
	rootCmd.Flags().BoolVarP(
		&ignoreUnstaged,
		"ignore-unstaged",
		"i",
		false,
		"ignore unstaged changes in the git diff",
	)
}
