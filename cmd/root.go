package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nox/packages"
	"github.com/spf13/cobra"
)

var version = "v0.0.1"
var key string
var configFile string

var rootCmd = &cobra.Command{
	Use:     "nox",
	Short:   "Nox: The CLI tool that brings sanity to your projects. Let’s get stuff done!",
	Long:    "Nox is your friendly neighborhood CLI tool that takes the pain out of managing your web projects. Focus on coding, let Nox handle the rest.",
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		if key != "" {
			packages.SetOpenAIKey(key)
		} else {
			fmt.Println("Welcome to Nox! I’m here to help you with your web projects. 🚀")
			fmt.Println("Type 'nox --help' to see what I can do for you.")
		}

	},
}

func init() {
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "Set OpenAI key")

	// Add config file for nox
	homeDir, _ := os.UserHomeDir()
	configFile = filepath.Join(homeDir, ".nox_config.json")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Uh-oh! Something went wrong: '%s'. Don't panic, just fix it and try again! 💻🚀", err)
		os.Exit(1)
	}
}
