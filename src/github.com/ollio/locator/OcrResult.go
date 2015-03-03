package main

import (
	"encoding/xml"
)

type Setup struct {
	Language string 	`xml:"language"`
	InputFile string	`xml:"inputfile"`
	Resolution string	`xml:"resolution"`
	MarketWidth string	`xml:"marketWidth"`
	FileTimeStamp string`xml:"filetimestamp"`
	OcrTime string		`xml:"ocrtime"`
}

type Location struct {
	System string		`xml:"system"`
	Station string		`xml:"station"`
	Player string		`xml:"player"`
}

type Market struct {
	Entries	[]Entry		`xml:"entry"`
}

type Entry struct {
	commodity string
	sell int16
	buy int16
	demand int16
	demandlevel string
	supply int16
	supplylevel string
}

type OcrResult struct {
	XMLName xml.Name 	`xml:"ocrresult"`
	Setup Setup			`xml:"setup"`
	Location Location	`xml:"location"`
	Market Market		`xml:"market"`
}
