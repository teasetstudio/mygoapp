package config

import "os"

var HomeDir, _ = os.UserHomeDir()
var AppDataDir = HomeDir + "\\AppData\\Local"
var AppDir = AppDataDir + "\\mygoapp"
var DownloadDir = HomeDir + "\\Downloads"
