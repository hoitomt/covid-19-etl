package main

import (
	"database/sql"
	"strings"
	"time"
)

type County struct {
	Id        int       `db:"id"`         // id INT PRIMARY KEY
	Name      string    `db:"name"`       // name VARCHAR
	Fips      string    `db:"fips"`       // fips VARCHAR
	StateId   int       `db:"state_id"`   // state_id INT
	CreatedAt time.Time `db:"created_at"` // created_at timestamp without time zone NOT NULL
	UpdatedAt time.Time `db:"updated_at"` // updated_at timestamp without time zone NOT NULL
}

func (db *CovidDb) GetCountyByName(countyName string) *County {
	var err error

	county := County{}
	err = db.Get(&county, "SELECT * FROM counties WHERE lower(name) = $1", strings.ToLower(countyName))

	switch err {
	case sql.ErrNoRows:
		logger.Infof("No rows returned from State query. Do not insert until state has been inserted.")
		return nil
	case nil:
		return &county
	default:
		logger.Fatalf("Bad things man - Select Id2 %s", err)
		return nil
	}
}
