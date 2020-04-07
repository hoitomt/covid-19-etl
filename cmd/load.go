package main

import (
	"log"
	"sync"
)

func Load(loadChan chan CovidCase, wg *sync.WaitGroup) {
	for covidCase := range loadChan {
		log.Printf("Load CovidCase %#v", covidCase.ToCaseSql())
	}
	log.Println("Data Load Complete")
	wg.Done()
}
