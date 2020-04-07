package main

import "log"

func Load(loadChan chan CovidCase) {
	for covidCase := range loadChan {
		log.Printf("Load CovidCase %#v", covidCase.ToCaseSql())
	}

}
