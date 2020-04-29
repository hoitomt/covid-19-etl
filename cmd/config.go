package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	DbUserName string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

var Config ViperConfig

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

	log.Println("Set Config variables")
	Config.DbHost = viper.GetString("database.host")
	Config.DbPort = viper.GetString("database.port")
	Config.DbUserName = viper.GetString("database.user_name")
	Config.DbPassword = viper.GetString("database.password")
	Config.DbName = viper.GetString("database.database")
}
