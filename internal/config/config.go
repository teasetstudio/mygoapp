package config

import (
	"log"
	"os"
	"runtime"
)

var HomeDir, _ = os.UserHomeDir()
var UserConfigDir, _ = os.UserConfigDir()
var AppDir = UserConfigDir + "\\mygoapp"
var InvoiceDataFileName = "invoice_data.yaml"
var InvoiceDataFile = AppDir + "\\" + InvoiceDataFileName
var DownloadDir = HomeDir + "\\Downloads"
var BusinessDir = "C:\\busines"
var FontsDir = "C:\\Windows\\Fonts"

func init() {
	appConfigFolderSetup()

	exists, _ := FileExists(InvoiceDataFile)
	if exists {
		ReadYAMLInvoiceConfigFile()
	} else {
		SetYAMLInvoiceDataFile(DefaultInvoiceData)
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
