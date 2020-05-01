package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CovidDb struct {
	*sqlx.DB
}

func NewDB() CovidDb {
	logger.Info("Initialize database")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		Config.DbUserName, Config.DbPassword, Config.DbName, Config.DbHost, Config.DbPort)

	logger.Infof("Connect to Database: %s", connStr)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logger.Fatalf("Failed to connect to the database %s", err)
	}

	database := CovidDb{db}
	return database
}
