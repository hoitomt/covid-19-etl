package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func Transform(transformChan chan CaseFile, loadChan chan CovidCase) {
	log.Println("Start Transform")
	for caseFile := range transformChan {
		log.Printf("Start processing case file %s, %s", caseFile.FileType, caseFile.FilePath)

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
				if line[0] == "date" {
					continue
				}
				if caseFile.FileType == "county" {
					loadChan <- NewCountyCase(line)
				} else if caseFile.FileType == "state" {
					loadChan <- NewStateCase(line)
				}
			}
		}
		log.Printf("Complete processing case file %s, %s", caseFile.FileType, caseFile.FilePath)
	}
	close(loadChan)
}
