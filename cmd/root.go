package cmd

import (
	"fmt"
	"os"

	"github.com/aethiopicuschan/cmg/cmd/config"
	pkgConfig "github.com/aethiopicuschan/cmg/pkg/config"
	"github.com/aethiopicuschan/cmg/pkg/git"
	"github.com/aethiopicuschan/cmg/pkg/llm"
	"github.com/aethiopicuschan/cmg/pkg/logs"
	"github.com/spf13/cobra"
)

var details bool

var rootCmd = &cobra.Command{
	Use:          "cmg",
	Long:         `Commit message generator based on git diff using an LLM`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		cfg, err := pkgConfig.GetConfig()
		if err != nil {
			logs.Fatal(err.Error())
		}

		// Ensure chat model config exists
		if cfg.LLM.Models.Chat == nil {
			logs.Fatal("chat model configuration is required")
		}

		// Create LLM (ToolCallingChatModel)
		chatModel, err := llm.NewChatModel(cfg.LLM)
		if err != nil {
			logs.Fatal(err.Error())
		}

		// Build diff options (LLM-safe defaults)
		diffOpts := git.DiffOptions{
			MaxTotalBytes:   120_000,
			MaxPerFileBytes: 8_000,
			IncludeDiffBody: true,
		}

		ctx := cmd.Context()

		// Generate commit message
		message, err := llm.GenerateCommitMessage(
			ctx,
			chatModel,
			diffOpts,
			details,
		)
		if err != nil {
			logs.Fatal(err.Error())
		}

		// Output result
		fmt.Fprintln(cmd.OutOrStdout(), message)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(config.Cmd)
	rootCmd.Flags().BoolVar(
		&details,
		"details",
		false,
		"include detailed commit body (multi-line commit message)",
	)
}
