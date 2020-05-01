package main

import (
	"sync"
)

type CaseFile struct {
	FilePath string
	FileType string
	DataChan chan CovidCase
}

func main() {
	initConfig()
	initLogger()

	logger.Println("Start Application")
	var wg sync.WaitGroup

	logger.Info("Step")
	fileTypes := []string{"county", "state"}
	transformChan := make(chan CaseFile)
	loadChan := make(chan CovidCase)

	wg.Add(1)
	go Load(loadChan, &wg)
	go Transform(transformChan, loadChan)
	// wg.Wait()
	// os.Exit(0)

	wg.Add(1) // done is at the end of the Load process
	for _, fileType := range fileTypes {
		fileToTransform, err := Extract(fileType)
		if err != nil {
			logger.Errorf("Error downloading %s data. %s", fileType, err)
			continue
		}

		caseFile := CaseFile{fileToTransform, fileType, make(chan CovidCase)}
		transformChan <- caseFile
	}
	close(transformChan)
	wg.Wait()
}
