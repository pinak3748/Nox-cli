package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/nox/packages"
	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var pageName string

var rootCmd = &cobra.Command{
	Use:     "nox",
	Short:   "Nox: The CLI tool that brings sanity to your projects. Let’s get stuff done!",
	Long:    "Nox is your friendly neighborhood CLI tool that takes the pain out of managing your web projects. Focus on coding, let Nox handle the rest.",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if pageName != "" {
			packages.Page(strings.ToLower(pageName))
		} else {
			packages.Page(strings.ToLower(pageName))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Uh-oh! Something went wrong: '%s'. Don't panic, just fix it and try again! 💻🚀", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&pageName, "page", "p", "", "Tell me the page you want to create, and I’ll whip it up in no time!")
}
