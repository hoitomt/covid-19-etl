package main

import (
	"log"
	"sync"
)

type CaseFile struct {
	FilePath string
	FileType string
	DataChan chan CovidCase
}

func main() {
	log.Println("Start Application")
	initConfig()
	var wg sync.WaitGroup

	fileTypes := []string{"county", "state"}

	transformChan := make(chan CaseFile)
	loadChan := make(chan CovidCase)

	go Load(loadChan, &wg)
	go Transform(transformChan, loadChan)

	wg.Add(1) // done is at the end of the Load process
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
	wg.Wait()
}
