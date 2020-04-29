package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLogging() *logrus.Logger {
	logger := logrus.New()
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
	return logger
}
