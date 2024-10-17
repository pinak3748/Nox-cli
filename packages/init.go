package packages

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
)

type ProjectType string
type PackageManager string

const (
	React_JS ProjectType    = "react + js"
	React_TS ProjectType    = "react + ts"
	NPM      PackageManager = "npm"
	Yarn     PackageManager = "yarn"
	Bun      PackageManager = "bun"
	Pnpm     PackageManager = "pnpm"
)

type Project struct {
	Name           string
	ProjectType    ProjectType
	PackageManager PackageManager
}

func Init(projectName string) {
	fmt.Println("✨ Creating a new project directory... Let’s get this party started! 🎉")

	if projectName == "" {
		projectName = "client"
	}

	var basePath = projectName + "/"
	var projectType ProjectType
	var packageManager PackageManager

	// Ask the user to choose a project type
	projectTypePrompt := promptui.Select{
		Label: "🤔 What type of project would you like to create?",
		Items: []string{"React + JavaScript", "React + TypeScript"},
	}

	_, projectTypeChoice, err := projectTypePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch projectTypeChoice {
	case "React + JavaScript":
		projectType = React_JS
	case "React + TypeScript":
		projectType = React_TS
	default:
		fmt.Println("⚠️ Invalid choice. Defaulting to React + JavaScript.")
		projectType = React_JS
	}

	fmt.Println("🚧 Creating a new", projectType, "project...")

	// Ask the user to choose a package manager
	packageManagerPrompt := promptui.Select{
		Label: "🤔 Which package manager would you like to use?",
		Items: []string{"npm", "yarn", "bun", "pnpm"},
	}

	_, packageManagerChoice, err := packageManagerPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch packageManagerChoice {
	case "npm":
		packageManager = NPM
	case "yarn":
		packageManager = Yarn
	case "bun":
		packageManager = Bun
	case "pnpm":
		packageManager = Pnpm
	default:
		fmt.Println("⚠️ Invalid choice. Defaulting to npm.")
		packageManager = NPM
	}

	// Clone a repository
	var githubRepo = "https://github.com/pinak3748/Seed-UI.git"

	// Check if the "client" directory exists at the root level
	if _, err := os.Stat("client"); os.IsNotExist(err) {

		var gitCloneCmd = exec.Command("git", "clone", githubRepo, projectName)

		err = gitCloneCmd.Run()
		if err != nil {
			fmt.Println("⚠️ Oops! Something went wrong while cloning the repository. Check your connection and try again.")
			os.Exit(1)
		}
	}

	// Installing dependencies
	var installCmd *exec.Cmd

	switch packageManager {
	case NPM:
		installCmd = exec.Command("npm", "install")
	case Yarn:
		installCmd = exec.Command("yarn", "install")
	case Bun:
		installCmd = exec.Command("bun", "install")
	case Pnpm:
		installCmd = exec.Command("pnpm", "install")
	}

	installCmd.Dir = basePath

	// Start the spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	var Suffixes = []string{
		"🚚 Installing dependencies...",
		"🔄 Setting things up...",
		"📦 Fetching packages...",
		"⚙️ Configuring project...",
		"🔧 Building modules...",
		"🛠️ Preparing environment...",
		"⏳ This might take a minute, hang tight...",
		"🙌 Almost there, just a moment...",
		"🔍 Checking everything...",
		"🚀 Finalizing setup...",
	}
	go func() {
		for {
			for _, suffix := range Suffixes {
				s.Suffix = " " + suffix
				time.Sleep(1 * time.Second)
			}
		}
	}()
	s.Start()

	err = installCmd.Run()

	// Stop the spinner
	s.Stop()

	if err != nil {
		fmt.Println("⚠️ Oops! Something went wrong while installing dependencies. Please try again.")
		os.Exit(1)
	}

	fmt.Println("🚀 Setup completed successfully! You're ready to roll. Time to build something awesome! 💻✨")
}
