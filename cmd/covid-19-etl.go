package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func initConfig() {
	pflag.StringP("config_file_path", "c", "configs/config.yml", "Path to the configuration file")

	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	log.Printf("Config File Path %s", viper.GetString("config_file_path"))
}

func main() {
	log.Println("Start Application")
	initConfig()
	Extract()

}
