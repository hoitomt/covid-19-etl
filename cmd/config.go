package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

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
	env := os.Getenv("ENV")
	fmt.Printf("Environment: %s\n", env)

	basePath := projectBasePath()
	fmt.Printf("Project Base Path: %s\n", basePath)

	configFilePath := pflag.StringP("config_file_path",
		"c",
		filepath.Join(basePath, "config"),
		"Path to the configuration file")
	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading the configuration file %s. %s", *configFilePath, err)
	}

	log.Println("Set Config variables")
	Config.DataCountyBasePath = viper.GetString(fmt.Sprintf("%s.data.county.base_path", env))
	Config.DataCountyUrl = viper.GetString(fmt.Sprintf("%s.data.county.url", env))
	Config.DataStateBasePath = viper.GetString(fmt.Sprintf("%s.data.state.base_path", env))
	Config.DataStateUrl = viper.GetString(fmt.Sprintf("%s.data.state.url", env))

	Config.DbHost = viper.GetString(fmt.Sprintf("%s.database.host", env))
	Config.DbPort = viper.GetString(fmt.Sprintf("%s.database.port", env))
	Config.DbUserName = viper.GetString(fmt.Sprintf("%s.database.user_name", env))
	Config.DbPassword = viper.GetString(fmt.Sprintf("%s.database.password", env))
	Config.DbName = viper.GetString(fmt.Sprintf("%s.database.name", env))

	Config.LogFilePath = filepath.Join(basePath, viper.GetString(fmt.Sprintf("%s.log.path", env)))
	Config.LogLevel = viper.GetString(fmt.Sprintf("%s.log.level", env))
}

func projectBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
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
