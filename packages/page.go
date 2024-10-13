package packages

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nox/content"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Page(pageName string) {
	fmt.Println("✨ Creating a new page... Hold tight, magic is happening! 📄✨")

	titleCaser := cases.Title(language.English)

	// Check if the "client" directory exists at the root level
	if _, err := os.Stat("client"); os.IsNotExist(err) {
		fmt.Println("⚠️ Uh-oh! Looks like you're missing the 'client' directory. Run this command with the correct directory flag! 🏗️")
		os.Exit(1)
	}

	// Create the new directory under "pages" for the specified page
	createDir(filepath.Join("client", "src", "pages", pageName))

	var basePath = filepath.Join("client", "src", "pages")
	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		fmt.Println("😬 Sorry, the 'pages' directory doesn’t exist, or it’s already there: ", err)
		os.Exit(1)
	}

	// Create the "modules" directory for the new page
	createDir(filepath.Join(basePath, pageName, "modules"))

	var modulePath = filepath.Join(basePath, pageName, "modules")
	createPage(filepath.Join(modulePath, "index.ts"), content.GenerateMainPageContent(titleCaser.String(pageName)))

	// Create the Redux Toolkit files for the new page
	createDir(filepath.Join(basePath, pageName, "features"))

	var pageActionPath = filepath.Join(basePath, pageName, "features", pageName+"Action.ts")
	var pageSlicePath = filepath.Join(basePath, pageName, "features", pageName+"Slice.ts")

	createPage(pageActionPath, content.GenerateActionsContent(titleCaser.String(pageName)))
	createPage(pageSlicePath, content.GenerateSliceContent(titleCaser.String(pageName)))

	// Create the index file for the new page
	createPage(filepath.Join(basePath, pageName, "index.ts"), content.GenerateIndexContent(titleCaser.String(pageName)))

	fmt.Println("🎉 Success! Your new page is live and ready to go. Time to get creative! 🚀")
}

func createPage(pagePath string, content string) {
	file, err := os.Create(pagePath)
	if err != nil {
		fmt.Println("💥 Yikes! Something went wrong while creating the file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("⚠️ Oops! There was an issue writing content to the file.")
		os.Exit(1)
	}
}

func createDir(dirPath string) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("💥 Oops, we hit a snag while creating the directory: ", err)
		os.Exit(1)
	}
}
