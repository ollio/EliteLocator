package main

import (
	"os"
	"log"
	"time"
	"path/filepath"
)

const VERSION = "v1.02"

func main() {
	log.Println("Elite Locator " + VERSION)

	appPath := appPath()
	patchAppConfig(appPath)

	logPath := logPath(appPath)

	player := new(Player)

	for {
		var updated = GetPlayer(logPath)

		if updated != nil && !equals(player, updated) && len(updated.Name) > 0 {
			player = updated
			PostPlayer(player)
			log.Println("update sent for "+player.Name + " -> " + player.System)
//			log.Println("Entry Date: " + player.UserData.EntryDate)
//			log.Println("Entry Data: " + player.UserData.Data)
//			log.Println("Entry Data: " + player.LogDate)
		}

		time.Sleep(30*time.Second)
	}
}

func appPath() string {
	// Log Path Override
	if len(os.Args) >= 2 {
		return os.Args[1]
	}

	myDir := filepath.ToSlash(filepath.Dir(os.Args[0]))

	// Check if we are in the target Folder
	ok, _ := filepath.Match("*/Products/FORC-FDEV*", myDir)
	if ok {
		return filepath.Clean(myDir)
	}

	match, myDir := find(myDir)
	if match {
		return filepath.Clean(myDir)
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
