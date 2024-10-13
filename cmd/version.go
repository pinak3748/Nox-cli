package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Find out which version of Nox you’re running",
	Long:  `All great software evolves. Here’s the current flavor of Nox you're using!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔥 You're running Nox v0.1 — stay awesome and keep coding! 🚀")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
