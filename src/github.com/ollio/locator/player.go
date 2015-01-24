package main

type Player struct {
	Name string      	`json:"name"`
	System string       `json:"system"`
	LastUpdate string   `json:"lastUpdate"`
	Online bool         `json:"online"`
	Health float64      `json:"health"`
}


