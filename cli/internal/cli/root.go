package cli

import (
	"github.com/spf13/cobra"
	"github.com/your-org/your-cli/internal/commands"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Cosmos CLI template",
	Long:  "A CLI template with plug-and-play commands. Add subcommands in internal/commands or pkg/cli/commands.",
}

// Execute runs the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(commands.VersionCmd)
	rootCmd.AddCommand(commands.RunCmd)
	// Plug more commands: rootCmd.AddCommand(commands.YourCmd) or from pkg/cli/commands
}
