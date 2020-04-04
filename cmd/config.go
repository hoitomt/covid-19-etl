package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func initConfig() {
	configFilePath := pflag.StringP("config_file_path", "c", "configs", "Path to the configuration file")

	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading the configuration file %s. %s", *configFilePath, err)
	}

	logFilePath := viper.GetString("log.base_path")
	log.Printf("Create Log Path: %v", logFilePath)
	err = os.MkdirAll(logFilePath, 0755)
	if err != nil {
		log.Fatalf("Error creating the log directory %s. %s", logFilePath, err)
	}
}
