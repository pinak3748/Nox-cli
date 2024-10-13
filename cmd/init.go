package cmd

import (
	"github.com/nox/packages"
	"github.com/spf13/cobra"
)

var projectName string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Kickstart a brand-new project with style",
	Long:  `Need a fresh start? "init" gets you rolling with a shiny new project, ready for action!`,
	Run: func(cmd *cobra.Command, args []string) {
		packages.Init(projectName)
	},
}

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Give your project a name, because every masterpiece deserves a title!")
	rootCmd.AddCommand(initCmd)
}
