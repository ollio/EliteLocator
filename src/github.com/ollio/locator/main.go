package main

import (
	"os"
	"log"
	"time"
)

func main() {
	log.Println("Elite Locator")

	logpath := logPath()

	for {
		var player = GetPlayer(logpath)
		PostPlayer(player)

		time.Sleep(1*time.Minute)
	}
}

func logPath() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "logs"
}
