package main

import (
	"log"

	"github.com/hoitomt/covid-19-etl-go/etl"
)

func main() {
	log.Println("Start Application")
	etl.Extract()

}
