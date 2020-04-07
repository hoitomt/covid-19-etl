package main

import (
	"log"
)

type CaseFile struct {
	FilePath string
	FileType string
	DataChan chan CovidCase
}

func main() {
	log.Println("Start Application")
	initConfig()

	fileTypes := []string{"county", "state"}

	transformChan := make(chan CaseFile)
	loadChan := make(chan CovidCase)

	go Load(loadChan)
	go Transform(transformChan, loadChan)

	for _, fileType := range fileTypes {
		fileToTransform, err := Extract(fileType)
		if err != nil {
			log.Printf("Error downloading %s data. %s", fileType, err)
			continue
		}

		caseFile := CaseFile{fileToTransform, fileType, make(chan CovidCase)}
		transformChan <- caseFile
	}
	close(transformChan)
}

// 1. Put the extracted file onto the transform channel
// 2. Loop through the file and put each record on the Load Channel
