package locator

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("Elite Locator \n")

	var player = GetPlayer(logPath())


	fmt.Printf("Player: " + player.Name + " in system: " + player.System)

}

func logPath() string {
	if len(os.Args) >= 2 {
		return os.Args[1]
	}
	return "logs"
}
