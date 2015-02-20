package main

type Player struct {
	Name string      	`json:"name"`
	System string       `json:"system"`
	Online bool         `json:"online"`
	Health float64      `json:"health"`
	Channel string 	    `json:"channel"`
	UserData UserData   `json:"userData"`
	LogDate string      `json:"logDate"`
	Version string		`json:"clientVersion"`
}


