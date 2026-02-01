package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runName string

var RunCmd = &cobra.Command{
	Use:   "run [args...]",
	Short: "Run a sample task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if runName == "" {
			runName = "world"
		}
		fmt.Printf("Hello, %s!\n", runName)
		return nil
	},
}

func init() {
	RunCmd.Flags().StringVarP(&runName, "name", "n", "", "name to greet")
}
