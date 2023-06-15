package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func setJSONConfigExample() {
	username := "JohnDoe"
	configData := []byte(fmt.Sprintf(`{
			"username": "%s",
			"key1": "value1",
			"key2": "value2",
			"key3": "value3"
		}`, username))

	SetJSONConfigFile("main", configData)
}

func SetJSONConfigFile(fileName string, configData []byte) {
	// Path to the config file
	filePath := filepath.Join(AppDir, fileName+"_config.json")

	// Create the config file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create the config file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write the config data to the file
	_, err = file.Write(configData)
	if err != nil {
		fmt.Println("Failed to write the config data:", err)
		os.Exit(1)
	}

	fmt.Println("Config file created successfully: ", filePath)
}
