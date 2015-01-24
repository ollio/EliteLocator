package main

import (
	"log"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

func PostPlayer(player *Player) {
	url := "http://elitelocator.herokuapp.com/locate/player/"
//	log.Println("URL: ", url)

	json, _ := json.Marshal(player)
	log.Println("json: " + string(json))

	request := gorequest.New()
	_, _, errs := request.Post(url).
	Set("Content-Type","application/json").
	Send(string(json)).
	End()

	if errs != nil {
		log.Println(errs)
	}
}


