package locator

import (
	"fmt"
	"io/ioutil"
)

func GetPlayer(path string) *Player {

	fmt.Println("Parsing " + path)

	files, _ := ioutil.ReadDir(path)

	for _, f := range files {

		fmt.Println(f.Name())
	}

	player := new(Player)
	player.Name = "Bozan"
	player.System = "Oluf"

	return player
}

