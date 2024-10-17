package packages

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	OpenAIKey string `json:"openai_key"`
}

var configFile string

func SetOpenAIKey(key string) {
	fmt.Println("OPEN AI KEY", len(key))

	// Add config file for nox
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Something went wrong while getting the home directory.")
		return
	}
	configFile = filepath.Join(homeDir, "nox_config.json")

	// TODO: Need to add proper validation for the openAI key
	if len(key) <= 1 {
		fmt.Println("Please provide the API key")
		return
	}

	config := Config{OpenAIKey: key}

	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		fmt.Println("Something went wrong while saving the key.")
		return
	}

	err = os.WriteFile(configFile, data, 0600)
	if err != nil {
		fmt.Println("Something went wrong while writing the config file.")
		return
	}

	fmt.Println("Great! OpenAI key saved successfully")
}
