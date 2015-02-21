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

func patchAppConfig(path string) {
	appConfig := path + "/AppConfig.xml"

	b, err := ioutil.ReadFile(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	reNetwork := regexp.MustCompile("(<Network[\\s\\S]*?<\\/Network>)")
	index := reNetwork.FindStringSubmatchIndex(string(b))

	if len(index) <=1 {
		reNetwork := regexp.MustCompile("(<Network[\\s\\S]*?\\/>)")
		index = reNetwork.FindStringSubmatchIndex(string(b))

		if len(index) <=1 {
			log.Fatal("AppConfig.xml is wrong!")
		}
	}

	begin:=b[:index[0]]
	end:=b[index[1]:]

/*
	log.Println("Begin: " + string(begin))
	log.Println("End: " + string(end))
*/

	// open output file
	fo, err := os.Create(appConfig)
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

