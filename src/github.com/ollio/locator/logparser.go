package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"bufio"
	"path/filepath"
	"regexp"
	"strings"
)

type LogfileData struct {
	LogDate    time.Time
	PlayerName string
	System     string
}

func GetPlayer(path string) *Player {

	fmt.Println("Parsing " + path)

	/*
		files, _ := ioutil.ReadDir(path)

		for _, f := range files {
			fmt.Println(f.Name())
		}
	*/

	files, _ := filepath.Glob(path + "/netLog.*.log")

	/*
		for _, f := range files {
			fmt.Println("File: ", f)
		}
	*/

	var logfileData = parse(files[len(files) - 1])
	//	fmt.Println("Last: ", files[len(files)-1])


	player := new(Player)
	player.Name = logfileData.PlayerName
	player.System = logfileData.System

	return player
}

func parse(fileName string) *LogfileData {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumeber uint32
	ld := new(LogfileData)
	reSystem := regexp.MustCompile("^.*\\sSystem:[0-9].(.*)\\)\\sBody.*$")
	reFindBestIsland := regexp.MustCompile("^.*\\sFindBestIsland:(.*):.*$")

	for scanner.Scan() {
		//		fmt.Printf("%d : %s \n", lineNumeber, scanner.Text())

		line := scanner.Text()
		if lineNumeber == 0 {
			ld.LogDate = getDate(line)
		} else if strings.Contains(line, "System") {
			match := reSystem.FindStringSubmatch(line)
			if len(match) > 1 {
				ld.System = match[1]
			}
		} else if strings.Contains(line, "FindBestIsland") {
			match := reFindBestIsland.FindStringSubmatch(line)
			if len(match) > 1 {
				ld.PlayerName = match[1]
			}
		}

		lineNumeber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ld
}

func getDate(line string) time.Time {
	const layout = "06-01-02-15:04"
	date, err := time.Parse(layout, line[:14])
	if err != nil {
		log.Fatal(err)
	}
	return date
}



