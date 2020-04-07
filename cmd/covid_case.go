package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type CovidCase interface {
	ToCaseSql() string
	ToEntitySql() string
}

type CountyCase struct {
	Date   *time.Time
	County string
	State  string
	Fips   string
	Cases  int
	Deaths int
}

func NewCountyCase(csvString []string) CountyCase {
	countyCase := CountyCase{}
	countyCase.parseCsv(csvString)
	return countyCase
}

func (c *CountyCase) parseCsv(csvString []string) {
	dateVal, err := time.Parse("2006-01-02", csvString[0])
	c.Date = &dateVal
	if err != nil {
		// c.Date = nil
	}
	c.County = csvString[1]
	c.State = csvString[2]
	c.Fips = csvString[3]
	c.Cases, err = strconv.Atoi(csvString[4])
	if err != nil {
		log.Printf("ERROR parsing case value %s", csvString[3])
		c.Cases = 0
	}
	c.Deaths, err = strconv.Atoi(csvString[5])
	if err != nil {
		log.Printf("ERROR parsing death value %s", csvString[4])
		c.Deaths = 0
	}
}

func (c CountyCase) ToCaseSql() string {
	return fmt.Sprintf("%s, %s, %s", c.Date, c.County, c.State)
	// return "", nil
}

func (c CountyCase) ToEntitySql() string {
	return ""
}

////////// State Case /////////////

type StateCase struct {
	Date   *time.Time
	County string
	State  string
	Fips   string
	Cases  int
	Deaths int
}

func NewStateCase(csvString []string) StateCase {
	stateCase := StateCase{}
	stateCase.parseCsv(csvString)
	return stateCase
}

func (stateCase *StateCase) parseCsv(csvString []string) {
	var err error

	dateVal, err := time.Parse("2006-01-02", csvString[0])
	stateCase.Date = &dateVal
	if err != nil {
		stateCase.Date = nil
	}
	stateCase.State = csvString[1]
	stateCase.Fips = csvString[2]
	stateCase.Cases, err = strconv.Atoi(csvString[3])
	if err != nil {
		log.Printf("ERROR parsing case value %s", csvString[3])
		stateCase.Cases = 0
	}
	stateCase.Deaths, err = strconv.Atoi(csvString[4])
	if err != nil {
		log.Printf("ERROR parsing death value %s", csvString[4])
		stateCase.Deaths = 0
	}
}

func (stateCase StateCase) ToCaseSql() string {
	return fmt.Sprintf("%s, %s", stateCase.Date, stateCase.State)
	// return "", nil
}

func (stateCase StateCase) ToEntitySql() string {
	return ""
}
