package config

import (
	"log"
	"os"
	"runtime"
)

var HomeDir, _ = os.UserHomeDir()
var UserConfigDir, _ = os.UserConfigDir()
var AppDir = UserConfigDir + "\\mygoapp"
var ConfigFileName = "config.yaml"
var ConfigFile = AppDir + "\\" + ConfigFileName
var DownloadDir = HomeDir + "\\Downloads"
var BusinessDir = "C:\\busines"

func init() {
	appConfigFolderSetup()

	exists, _ := FileExists(ConfigFile)
	if exists {
		readYAMLInvoiceConfigFile()
	} else {
		setYAMLInvoiceConfigFile(DefaultInvoiceConfig)
	}
}

func appConfigFolderSetup() {
	if runtime.GOOS == "windows" {
		if err := os.MkdirAll(AppDir, os.ModePerm); err != nil {
			if os.IsExist(err) {
				// check that the existing path is a directory
				info, err := os.Stat(AppDir)
				if err != nil {
					log.Fatal(err)
				}
				if !info.IsDir() {
					log.Fatal("path exists but is not a directory")
				}
			} else {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("OS not supported")
	}
}
