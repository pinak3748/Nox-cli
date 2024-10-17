package packages

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetOpenAIKey() (string, error) {
	// Add config file for nox
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Something went wrong while getting the home directory.")
		return "", err
	}
	configFile = filepath.Join(homeDir, "nox_config.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config.OpenAIKey, nil
}
