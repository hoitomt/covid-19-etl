package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

func Load(loadChan chan CovidCase, wg *sync.WaitGroup) {
	logger.Info("Start Load")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		Config.DbUserName, Config.DbPassword, Config.DbName, Config.DbHost, Config.DbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatalf("Failed to connect to the database %s", err)
	}

	logger.Infof("Connection String: %s", connStr)
	for covidCase := range loadChan {
		logger.Infof("Load CovidCase %#v", covidCase.ToCaseSql())
		covidCase.Upsert(db)
	}
	logger.Info("Data Load Complete")
	wg.Done()
}
