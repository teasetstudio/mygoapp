package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func SetYAMLFile(fileName string, data interface{}) error {
	configFilePath := filepath.Join(AppDir, fileName)

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal struct to YAML: %v", err)
	}

	// Write the new YAML data to the file
	err = os.WriteFile(configFilePath, yamlData, 0755)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}

	fmt.Println("Config file created successfully.")
	return nil
}

func ReadYAMLFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return data, nil
}

func setYAMLInvoiceConfigFile(config InvoiceConfigType) {
	err := SetYAMLFile(ConfigFileName, config)
	if err != nil {
		log.Fatalf("failed to set YAML config file: %v", err)
	}
}

func readYAMLInvoiceConfigFile() {
	data, err := ReadYAMLFile(ConfigFile)
	if err != nil {
		log.Fatalf("failed to read YAML config file: %v", err)
	}

	invoiceConfig := InvoiceConfigType{}
	err = yaml.Unmarshal(data, &invoiceConfig)
	if err != nil {
		log.Fatal(err)
	}
	InvoiceConfig = invoiceConfig
}
