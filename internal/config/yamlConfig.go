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

func setYAMLInvoiceDataFile(data InvoiceDataType) {
	err := SetYAMLFile(InvoiceDataFileName, data)
	if err != nil {
		log.Fatalf("failed to set YAML config file: %v", err)
	}
}

func readYAMLInvoiceConfigFile() {
	data, err := ReadYAMLFile(InvoiceDataFile)
	if err != nil {
		log.Fatalf("failed to read YAML config file: %v", err)
	}

	invoiceData := InvoiceDataType{}
	err = yaml.Unmarshal(data, &invoiceData)
	if err != nil {
		log.Fatal(err)
	}
	InvoiceData = invoiceData
}
