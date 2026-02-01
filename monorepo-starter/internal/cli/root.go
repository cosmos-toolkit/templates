package cli

import (
	"github.com/spf13/cobra"
	"github.com/your-org/your-app/internal/commands"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Monorepo starter CLI",
	Long:  "API + Worker + CLI monorepo. Use subcommands: version, run.",
}

// Execute runs the root command.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(commands.VersionCmd)
	rootCmd.AddCommand(commands.RunCmd)
}
