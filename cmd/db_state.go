package main

import (
	"database/sql"
	"strings"
	"time"
)

type State struct {
	Id        int       `db:"id"`         // id INT PRIMARY KEY
	Name      string    `db:"name"`       // name VARCHAR
	Fips      string    `db:"fips"`       // fips VARCHAR
	CreatedAt time.Time `db:"created_at"` // created_at timestamp without time zone NOT NULL
	UpdatedAt time.Time `db:"updated_at"` // updated_at timestamp without time zone NOT NULL
}

func (db *CovidDb) GetStateByName(stateName string) *State {
	var err error

	state := State{}
	err = db.Get(&state, "SELECT * FROM states WHERE lower(name) = $1", strings.ToLower(stateName))

	switch err {
	case sql.ErrNoRows:
		logger.Infof("No rows returned from State query. Do not insert until state has been inserted.")
		return nil
	case nil:
		return &state
	default:
		logger.Fatalf("Bad things man - Select Id2 %s", err)
		return nil
	}
}
