package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"bufio"
	"os"
)

const network =
	`<Network
		Port="0"
		upnpenabled="1"
		LogFile="netLog"
		DatestampLog="1"
		VerboseLogging="1"
	/>`

func patchAppConfig() {
	b, err := ioutil.ReadFile("AppConfig.xml")
	if err != nil {
		log.Fatal(err)
	}

	reNetwork := regexp.MustCompile("(<Network[\\s\\S]*?<\\/Network>)")

	index := reNetwork.FindStringSubmatchIndex(string(b))

	begin:=b[:index[0]]
	end:=b[index[1]:]

	// open output file
	fo, err := os.Create("AppConfig.xml")
	if err != nil {
		log.Fatal(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	w.Write(begin)
	w.WriteString(network)
	w.Write(end)

	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}


}

