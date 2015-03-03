package main

import (
	"os"
	"os/exec"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
	//	"bytes"
	"encoding/xml"
)

var checkedFiles = make([]string, 0)

func GetNewCommoditiesReport(player *Player) *OcrResult {

	files, _ := filepath.Glob(ScreenShotPath + "/*.bmp")

	for _, file := range files {
		if notIgnored(file) {
			//		log.Println(file)
			fi, _ := os.Stat(file)
			mtime := fi.ModTime()
			age := time.Now().Sub(mtime)

			if age.Hours() < 1 {
				return startScan(file, player)
			}
		}
	}

	return nil
}

func notIgnored(file string) bool {
	for _, f := range checkedFiles {
		if f == file {
			return false
		}
	}
	checkedFiles = append(checkedFiles, file)
	return true
}

func startScan(screenShot string, player *Player) *OcrResult {
	log.Println("Screenshot detected, start scan: " + screenShot)

	locatorDir, err := os.Getwd()
	err = os.Chdir(cfg.EliteOCR.Path)
	CheckError(err)

	copy(screenShot, "input.bmp")
	os.Remove("output.xml")

	err = os.Setenv("TESSDATA_PREFIX", filepath.Clean(cfg.EliteOCR.Path))
	CheckError(err)

	execBatch(filepath.Clean(cfg.EliteOCR.Path))

	output, err := ioutil.ReadFile("output.xml")
	CheckError(err)
//	log.Println(string(output))

	ocrResult := new(OcrResult)
	xml.Unmarshal(output, &ocrResult)

	if len(ocrResult.Location.System) == 0 {
		ocrResult.Location.System = player.System
	}

	ocrResult.Location.Player = player.Name

	log.Println("Scan success, System [" + ocrResult.Location.System + "] Station [" + ocrResult.Location.Station + "] Commodities:", len(ocrResult.Market.Entries))

	os.Chdir(locatorDir)
	CheckError(err)

	return ocrResult
}

func execBatch(tessDataEnv string) {
	batchFileName := "run.bat"

	os.Remove(batchFileName)
	batchFile, err := os.Create(batchFileName)
	CheckError(err)

	err = os.Setenv("TESSDATA_PREFIX", tessDataEnv)
	batchFile.WriteString("@echo off\n")
	batchFile.WriteString("set TESSDATA_PREFIX=\"" + tessDataEnv + "\" \n")
	batchFile.WriteString("setx -m TESSDATA_PREFIX \"" + tessDataEnv + "\" \n")
//	batchFile.WriteString("bin\\EliteOCRcmd.exe -i input.bmp -o output.xml -s \""+system+"\" \n")
	batchFile.WriteString("bin\\EliteOCRcmd.exe -i input.bmp -o output.xml \n")
	batchFile.Close()

	time.Sleep(1 * time.Second)

	_, err = exec.Command(batchFileName).Output()
	CheckError(err)
//	log.Println(string(out))

	os.Remove(batchFileName)
}

func copy(src string, dst string) {
	os.Remove(dst)

	r, err := os.Open(src)
	CheckError(err)
	defer r.Close()

	w, err := os.Create(dst)
	CheckError(err)
	defer w.Close()

	// do the actual work
	_, err = io.Copy(w, r)
	CheckError(err)
}
