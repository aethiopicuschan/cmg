package config

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "List current token holdings using Bubble Tea TUI",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation goes here
		// ...
	},
}
