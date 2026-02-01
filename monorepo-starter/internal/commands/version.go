package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	RunE:  func(cmd *cobra.Command, args []string) error { fmt.Println(version); return nil },
}
