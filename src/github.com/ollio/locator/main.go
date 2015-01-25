package main

import (
	"os"
	"log"
	"time"
)

func main() {
	log.Println("Elite Locator")

	logpath := logPath()

	player := new(Player)

	for {
		var updated = GetPlayer(logpath)

		if !compare(player, updated) {
			player = updated
			PostPlayer(player)
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

func compare(a, b *Player) bool {
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
	return true
}
