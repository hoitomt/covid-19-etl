package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.New()
	logger.Infof("Log file path %s", Config.LogFilePath)
	logFile, err := os.OpenFile(Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("error opening file: %v", err)
	}

	logger.Out = logFile

	if Config.LogLevel == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.Infof("Initialize Logging")
}
