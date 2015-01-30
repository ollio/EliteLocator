package main

import (
	"os"
	"log"
	"time"
	"bufio"
	"path/filepath"
	"regexp"
	"strings"
	"strconv"
)

func GetPlayer(path string) *Player {


	log.Println("Parsing " + path)

	files, _ := filepath.Glob(path + "/netLog.*.log")
	return parse(files[len(files) - 1])
	//	fmt.Println("Last: ", files[len(files)-1])
}

func parse(fileName string) *Player {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumeber uint32

	p := new(Player)
	p.Online = true

//	logDate := new(time.Time)

	reSystem := regexp.MustCompile("^.*\\sSystem:\\d+\\((.*)\\)\\sBody.*$")
	reFindBestIsland := regexp.MustCompile("^.*\\sFindBestIsland:(FRESH:)?(.*):.*$")
	reFindHealth := regexp.MustCompile("^.*\\shealth=([0-9\\.]*).*$")

	for scanner.Scan() {
		//		fmt.Printf("%d : %s \n", lineNumeber, scanner.Text())

		line := scanner.Text()
		if lineNumeber == 0 {
//			logDate := getDate(line)
		} else if strings.Contains(line, "System") {
			// "cruising" -> SuperCruse
			match := reSystem.FindStringSubmatch(line)
			if len(match) > 1 {
				p.System = match[1]
			}
		} else if strings.Contains(line, "FindBestIsland") {
			// FindBestIsland:FRESH:Bozan:18446744073709551615
			// "FRESH" -> Respawn
			// "id" -> current id
			match := reFindBestIsland.FindStringSubmatch(line)
			if len(match) > 2 {
				p.Name = match[2]
			}

			match2 := reFindHealth.FindStringSubmatch(line)
			if len(match2) > 1 {
				p.Health, err = strconv.ParseFloat(match2[1], 32)
			}
		} else if strings.Contains(line, "shutting down") {
			p.Online = false
		}

		lineNumeber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return p
}

func getDate(line string) time.Time {
	const layout = "06-01-02-15:04"
	date, err := time.Parse(layout, line[:14])
	if err != nil {
		log.Fatal(err)
	}
	return date
}



