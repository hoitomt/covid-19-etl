package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	DataCountyBasePath string
	DataCountyUrl      string
	DataStateBasePath  string
	DataStateUrl       string
	DbUserName         string
	DbPassword         string
	DbHost             string
	DbPort             string
	DbName             string
	LogFilePath        string
	LogLevel           string
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

	log.Println("Set Config variables")
	Config.DataCountyBasePath = viper.GetString("data.county.base_path")
	Config.DataCountyUrl = viper.GetString("data.county.url")
	Config.DataStateBasePath = viper.GetString("data.state.base_path")
	Config.DataStateUrl = viper.GetString("data.state.url")

	Config.DbHost = viper.GetString("database.host")
	Config.DbPort = viper.GetString("database.port")
	Config.DbUserName = viper.GetString("database.user_name")
	Config.DbPassword = viper.GetString("database.password")
	Config.DbName = viper.GetString("database.database")
	Config.LogFilePath = viper.GetString("log.path")
	Config.LogLevel = viper.GetString("log.level")
}

func (c *ViperConfig) DataBasePath(category string) string {
	if category == "county" {
		return c.DataCountyBasePath
	} else if category == "state" {
		return c.DataStateBasePath
	} else {
		return ""
	}
}

func (c *ViperConfig) DataUrl(category string) string {
	if category == "county" {
		return c.DataCountyUrl
	} else if category == "state" {
		return c.DataStateUrl
	} else {
		return ""
	}
}
