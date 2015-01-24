package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("Elite Locator \n")

	var player = GetPlayer(logPath())


	fmt.Println("Player: ", player.Name)
	fmt.Println("System: ", player.System)
	fmt.Println("Online: ", player.Online)
	fmt.Println("Health: ", player.Health)

}

func logPath() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "logs"
}
