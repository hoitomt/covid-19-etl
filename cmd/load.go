package main

import (
	"sync"

	_ "github.com/lib/pq"
)

func Load(loadChan chan CovidCase, wg *sync.WaitGroup) {
	logger.Info("Start Load")

	db := NewDB()
	for covidCase := range loadChan {
		logger.Infof("Load CovidCase %#v", covidCase.ToCaseSql())
		covidCase.Upsert(db)
	}
	logger.Info("Data Load Complete")
	wg.Done()
}
