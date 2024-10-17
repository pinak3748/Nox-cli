package packages

import (
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/nox/content"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Page(pageName string) {
	fmt.Println("✨ Creating a new page... Hold tight, magic is happening! 📄✨")

	titleCaser := cases.Title(language.English)
	capitalCaser := cases.Upper(language.English)
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
	createPage(filepath.Join(modulePath, "index.tsx"), content.GenerateMainPageContent(titleCaser.String(pageName)))

	// Create the Redux Toolkit files for the new page
	createDir(filepath.Join(basePath, pageName, "features"))

	var pageTypesPath = filepath.Join(basePath, pageName, "features", pageName+"Types.ts")
	var pageActionPath = filepath.Join(basePath, pageName, "features", pageName+"Action.ts")
	var pageSlicePath = filepath.Join(basePath, pageName, "features", pageName+"Slice.ts")

	createPage(pageActionPath, content.GenerateActionsContent(titleCaser.String(pageName)))
	createPage(pageSlicePath, content.GenerateSliceContent(titleCaser.String(pageName)))
	createPage(pageTypesPath, content.GenerateTypesContent(titleCaser.String(pageName)))
	// Create the index file for the new page
	createPage(filepath.Join(basePath, pageName, "index.ts"), content.GenerateIndexContent(titleCaser.String(pageName)))

	// Read the content of the routes.ts
	var constRoutesPath = filepath.Join("client", "src", "constants", "routes.ts")
	constRoutesContent, err := os.ReadFile(constRoutesPath)

	if err != nil {
		fmt.Println("⚠️ Oops! There was an issue reading the content of the routes file: ", err)
		os.Exit(1)
	}

	constRoutesContentString := string(constRoutesContent)
	newExportedRoute := fmt.Sprintf(" %s: '/%s',\n // New Page Route Goes Here!!! \n", capitalCaser.String(pageName), pageName)
	constRoutesContentString = strings.Replace(constRoutesContentString, "// New Page Route Goes Here!!!", newExportedRoute, 1)
	err = os.WriteFile(constRoutesPath, []byte(constRoutesContentString), 0644)
	if err != nil {
		fmt.Println("⚠️ Oops! There was an issue writing the updated content to the routes file: ", err)
		os.Exit(1)
	}

	// Read the content of the router.tsx
	var routesPath = filepath.Join("client", "src", "router.tsx")
	routesContent, err := os.ReadFile(routesPath)

	if err != nil {
		fmt.Println("⚠️ Oops! There was an issue reading the content of the routes file: ", err)
		os.Exit(1)
	}

	// Append the new route to the routes.js file
	routesContentString := string(routesContent)
	newRoute := fmt.Sprintf("  ,{\n    path: 'ROUTES.%s',\n    element: <%s />,\n  },\n // New Page Route Goes Here!!! \n", capitalCaser.String(pageName), titleCaser.String(pageName))
	newRoutesImport := fmt.Sprintf("import %s from './pages/%s';\n // New Page Route Import Goes Here!!! \n", titleCaser.String(pageName), pageName)
	routesContentString = strings.Replace(routesContentString, "// New Page Route Import Goes Here!!!", newRoutesImport, 1)

	routesContentString = strings.Replace(routesContentString, "// New Page Route Goes Here!!!", newRoute, 1)

	// Write the updated content back to the routes.js file
	err = os.WriteFile(routesPath, []byte(routesContentString), 0644)
	if err != nil {
		fmt.Println("⚠️ Oops! There was an issue writing the updated content to the routes file: ", err)
		os.Exit(1)
	}

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
