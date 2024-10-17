package cmd

import (
	"fmt"
	"os"

	"github.com/nox/packages"
	"github.com/spf13/cobra"
)

var name string
var description string

var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Create a new page",
	Long:  "Create a new page in your project",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Please provide a name for the page")
			os.Exit(1)
		}
		packages.Page(name)

		if description != "" {
			packages.DynamicComponent(name, description)
		}
	},
}

func init() {
	// name : name of the page
	pageCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the page to create")
	pageCmd.MarkFlagRequired("name")

	// description : Description of the UI you need for the component
	pageCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the UI you need for the component")

	rootCmd.AddCommand(pageCmd)

}
