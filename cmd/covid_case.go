package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/kisielk/sqlstruct"
)

type CovidCase interface {
	ToCaseSql() string
	ToEntitySql() string
	Upsert(*sql.DB)
}

type CountyCase struct {
	Date       *time.Time
	County     string
	State      string
	Fips       string
	Cases      int
	Deaths     int
	DbCountyId int
	DbCaseId   int
	DbStateId  int
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
		logger.Errorf("ERROR parsing case value %s", csvString[3])
		c.Cases = 0
	}
	c.Deaths, err = strconv.Atoi(csvString[5])
	if err != nil {
		logger.Errorf("ERROR parsing death value %s", csvString[4])
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

func (c CountyCase) Upsert(db *sql.DB) {
	// Determine if it is already in the database
	id := 0
	row := db.QueryRow("SELECT id FROM states WHERE name = '$1' LIMIT 1", c.State)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		logger.Infof("No rows returned from State query. %s", err)
	case nil:
		c.DbCountyId = id
	default:
		logger.Fatalf("Bad things man")
	}

	if c.DbCountyId > 0 {
		logger.Infof("County %s already exists", c.Fips)
	} else {
		// Fetch the state id
		row := db.QueryRow("SELECT id FROM states WHERE name = '$1' LIMIT 1", c.State)
		switch err := row.Scan(&id); err {
		case sql.ErrNoRows:
			logger.Println("No rows returned from State query. Do not insert until state has been inserted.")
			return
		case nil:
			c.DbStateId = id
		default:
			logger.Fatalf("Bad things man")
		}

		// Determine if the case is already present
		row = db.QueryRow("SELECT id FROM states WHERE name = '$1' LIMIT 1", c.State)
		switch err := row.Scan(&id); err {
		case sql.ErrNoRows:
			logger.Println("No rows returned from State query. Do not insert until state has been inserted.")
			return
		case nil:
			c.DbStateId = id
		default:
			logger.Fatalf("Bad things man")
		}

		loc, _ := time.LoadLocation("UTC")
		now := time.Now().In(loc)

		// Insert the county
		sqlStatement := `
		INSERT INTO counties (name, fips, state_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
		id := 0
		err := db.QueryRow(sqlStatement, c.County, c.DbStateId, now, now).Scan(&id)
		if err != nil {
			logger.Errorf("Error inserting county record %s", err)
		}
		c.DbCountyId = id

		sqlStatement = `
		INSERT INTO county_data (date, county_id, cases, deaths, sha256_hash, created_at, updated_at)
		VALUES ()
		`
	}
}

func (c CountyCase) sha256Hash() string {
	h := sha256.New()
	hashString := fmt.Sprintf("%s%s%s%s", c.Date, c.County, c.Cases, c.Deaths)
	h.Write([]byte(hashString))
	return fmt.Sprintf("%x", h.Sum(nil))
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
		logger.Errorf("ERROR parsing case value %s", csvString[3])
		stateCase.Cases = 0
	}
	stateCase.Deaths, err = strconv.Atoi(csvString[4])
	if err != nil {
		logger.Errorf("ERROR parsing death value %s", csvString[4])
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

func (stateCase StateCase) Upsert(db *sql.DB) {
}
