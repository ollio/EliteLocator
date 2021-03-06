package main

import (
	"log"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

func PostPlayer(player *Player) {
	url := "http://elitelocator.herokuapp.com/locate/player/"
//	url := "http://localhost:8080/locate/player/"
//	log.Println("URL: ", url)

	json, _ := json.Marshal(player)
//	log.Println("json: " + string(json))

	request := gorequest.New()
	_, _, errs := request.Post(url).
	Set("Content-Type","application/json").
	Send(string(json)).
	End()

	if errs != nil {
		log.Println(errs)
	}
}

func PostCommoditiesReport(commodities *OcrResult) {
	url := "http://elitelocator.herokuapp.com/commodities/"

	request := gorequest.New()
	_, _, errs := request.Post(url).
	Set("Content-Type","text/xml").
	// Evtl. Player Name / System
	Send(commodities).
	End()

	if errs != nil {
		log.Println(errs)
	}
}


