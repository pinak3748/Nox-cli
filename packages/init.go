package packages

import (
	"fmt"
	"os"
	"os/exec"
)

type Project struct {
	Name string
}

func Init(projectName string) {
	fmt.Println("✨ Creating a new project directory... Let’s get this party started! 🎉")

	if projectName == "" {
		projectName = "client"
	}

	// Clone a repository
	var githubRepo = "https://github.com/Ionio-io/seed-ui.git"

	var gitCloneCmd = exec.Command("git", "clone", githubRepo, projectName)

	err := gitCloneCmd.Run()
	if err != nil {
		fmt.Println("⚠️ Oops! Something went wrong while cloning the repository. Check your connection and try again.")
		os.Exit(1)
	}

	fmt.Println("🚀 Setup completed successfully! You're ready to roll. Time to build something awesome! 💻✨")
}
