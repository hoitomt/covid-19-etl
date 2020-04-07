package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Transform(transformChan chan CaseFile, loadChan chan CovidCase) {
	for caseFile := range transformChan {
		csvFile, err := os.Open(caseFile.FilePath)
		if err != nil {
			log.Printf("Error opening the csv file %s. %s", caseFile.FilePath, err)
			continue
		}
		reader := csv.NewReader(bufio.NewReader(csvFile))
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else {
				if caseFile.FileType == "county" {
					loadChan <- NewCountyCase(line)
				} else if caseFile.FileType == "state" {
					loadChan <- NewStateCase(line)
				}
			}
		}
	}
	close(loadChan)
}
