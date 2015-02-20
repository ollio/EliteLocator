package main

import (
	"os"
	"log"
	"bufio"
	"path/filepath"
	"regexp"
	"strings"
	"strconv"
)

func GetPlayer(path string) *Player {
	files, _ := filepath.Glob(path + "/netLog.*.log")
	if(len(files) > 0) {
		return parse(files[len(files) - 1])
	}
	log.Println("No Logfiles found !")
	return nil
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
	p.Version = VERSION

//	logDate := new(time.Time)

	reSystem := regexp.MustCompile("^.*\\sSystem:\\d+\\((.*)\\)\\sBody.*$")
	reFindBestIsland := regexp.MustCompile("^.*\\sFindBestIsland:(FRESH:)?(.*):(.*)$")
	reFindHealth := regexp.MustCompile("^.*\\shealth=([0-9\\.]*).*$")
	reEntryTime := regexp.MustCompile("^\\{(\\d{2}:\\d{2}:\\d{2})\\}\\s.*$")

	for scanner.Scan() {
		//		fmt.Printf("%d : %s \n", lineNumeber, scanner.Text())

		line := scanner.Text()
		if lineNumeber == 0 {
			p.LogDate = getDate(line)
		} else if strings.Contains(line, "System") {
			// "cruising" -> SuperCruse
			match := reSystem.FindStringSubmatch(line)
			if len(match) > 1 {
				p.System = match[1]
			}
		} else if strings.Contains(line, "FindBestIsland") {
			// FindBestIsland:FRESH:Bozan:18446744073709551615 << Channel of Player with ID
			// "FRESH" -> Respawn
			// "id" -> current id
			match := reFindBestIsland.FindStringSubmatch(line)
			switch len(match) {
			case 3:
				p.Name = match[2]
				break;
			case 4:
				p.Name = match[2]
				p.Channel = match[3]
				break;
			default:
				match2 := reFindHealth.FindStringSubmatch(line)
				if len(match2) > 1 {
					p.Health, err = strconv.ParseFloat(match2[1], 32)
				}
			}
		} else if strings.Contains(line, "shutting down") {
			p.Online = false
		} else if strings.Contains(line, "<data><users>") {
			p.UserData.Data = line
		} else if strings.Contains(line, "FriendsRequest") {
			match := reEntryTime.FindStringSubmatch(line)
			if len(match) > 1 {
				p.UserData.EntryTime = match[1]
			}
		}

		lineNumeber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return p
}

func getDate(line string) string {
	reLogDate := regexp.MustCompile("^(\\d{2}-\\d{2}-\\d{2}-\\d{2}:\\d{2})\\s.*$")
	match := reLogDate.FindStringSubmatch(line)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}



