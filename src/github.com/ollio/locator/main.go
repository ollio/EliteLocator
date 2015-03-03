package main

import (
	"os"
	"os/user"
	"log"
	"time"
	"path/filepath"
	"code.google.com/p/gcfg"
)

const VERSION = "v1.03"

var cfg Config
var ScreenShotPath string

func main() {
	log.Println("Elite Locator " + VERSION)

	err := gcfg.ReadFileInto(&cfg, configPath())
	CheckError(err)

	appPath := appPath()
	log.Println("appPath: " + appPath)

	logPath := logPath(appPath)
	ScreenShotPath = screenShotPath()
	log.Println("screenShotPath: " + ScreenShotPath)


	patchAppConfig(appPath)
	player := new(Player)

	for {
		var updated = GetPlayer(logPath)

		if updated != nil && !equals(player, updated) && len(updated.Name) > 0 {
			player = updated
			PostPlayer(player)
			log.Println("update sent for " + player.Name + " -> " + player.System)
		}

		if cfg.EliteOCR.Enable && len(player.System) > 0 {
			ocrResult := GetNewCommoditiesReport(player)
			if ocrResult != nil {
				PostCommoditiesReport(ocrResult);
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func configPath() string {
	// Config Path Override
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "locator.gcfg"
}

func appPath() string {
	var path string
	if len(cfg.EliteDangerous.Path) > 0 {
		path = cfg.EliteDangerous.Path
	} else {
		path = filepath.ToSlash(filepath.Dir(os.Args[0]))
	}

	// Recurse into subdirs and search the FORC-FDEV folder
	match, path := find(path)
	if match {
		return filepath.Clean(path)
	}

	log.Fatal("Can't find FORC-FDEV dir")
	return ""
}

func find(path string) (bool, string) {
	fileInfo, _ := os.Stat(path)
	if !fileInfo.IsDir() {
		return false, path
	}

	ok, _ := filepath.Match("*/Products/FORC-FDEV*", path)
	if ok {
		return true, path
	}

	files, _ := filepath.Glob(path + "/*")

	for _, f := range files {
		match, path := find(filepath.ToSlash(f))
		if match {
			return true, path
		}
	}

	return false, path
}

func logPath(appPath string) string {
	return filepath.Clean(appPath + "/logs/")
}

func screenShotPath() string {
	var screenShotPath string
	if len(cfg.EliteOCR.ScreenShotPath) > 0 {
		screenShotPath = cfg.EliteOCR.ScreenShotPath
	} else {
		usr, err := user.Current()
		CheckError(err)

		screenShotPath = usr.HomeDir+"/Pictures/Frontier Developments/Elite Dangerous"
	}
	return filepath.Clean(screenShotPath)
}

func equals(a, b *Player) bool {
	if &a == &b {
		return true
	}
	if a.Name != b.Name {
		return false
	}
	if a.System != b.System {
		return false
	}
	if a.Health != b.Health {
		return false
	}
	if a.Online != b.Online {
		return false
	}
	if a.Channel != b.Channel {
		return false
	}
	if a.UserData.EntryTime != b.UserData.EntryTime {
		return false
	}
	return true
}

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
