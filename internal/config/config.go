package config

import (
	"fmt"
	"os"
)

var HomeDir, _ = os.UserHomeDir()
var UserConfigDir, _ = os.UserConfigDir()
var AppDir = UserConfigDir + "\\mygoapp"
var ConfigFileName = "config.yaml"
var ConfigFile = AppDir + "\\" + ConfigFileName
var DownloadDir = HomeDir + "\\Downloads"
var BusinessDir = "C:\\busines"

func init() {
	err := os.MkdirAll(AppDir, 0755)
	if err != nil {
		fmt.Println("Failed to create the config file:", err)
		os.Exit(1)
	}

	exists, _ := FileExists(ConfigFile)
	if exists {
		readYAMLConfigFile()
	} else {
		setYAMLConfigFile(DefaultConfig)
	}
}
