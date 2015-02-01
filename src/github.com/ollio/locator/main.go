package main

import (
	"os"
	"log"
	"time"
)

func main() {
	log.Println("Elite Locator")

	patchAppConfig()

	logpath := logPath()

	player := new(Player)

	for {
		var updated = GetPlayer(logpath)

		if !equals(player, updated) && len(updated.Name) > 0 {
			player = updated
			PostPlayer(player)
			log.Println("update sent for "+player.Name + " -> " + player.System)
		}

		time.Sleep(30*time.Second)
	}
}

func logPath() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "logs"
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
	if a.Data != b.Data {
		return false
	}
	return true
}
