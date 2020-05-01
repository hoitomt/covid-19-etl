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

func (db *CovidDb) GetStateByName(stateName string) (*State, error) {
	var err error

	state := State{}
	err = db.Get(&state, "SELECT * FROM states WHERE lower(name) = $1", strings.ToLower(stateName))

	switch err {
	case sql.ErrNoRows:
		logger.Infof("No rows returned from State query. Do not insert until state has been inserted.")
		return nil, nil
	case nil:
		return &state, nil
	default:
		logger.Fatalf("Bad things man - Select Id2 %s", err)
		return nil, err
	}
}

func (db *CovidDb) InsertState(state State) error {
	var id int
	rows, err := db.NamedQuery("INSERT INTO states (name, fips) VALUES (:name, :fips) RETURNING id", state)
	if err != nil {
		logger.Errorf("ERROR inserting state: %v. %s", state, err)
		return err
	}
	if rows.Next() {
		rows.Scan(&id)
	}
	state.Id = id
	return nil
}

func (db *CovidDb) DeleteState(id int, stateName string) error {
	var err error
	if id > 0 {
		_, err = db.NamedQuery("DELETE FROM states where id = :id", id)
	} else {
		_, err = db.NamedQuery("DELETE FROM states where lower(name) = :name", stateName)
	}
	if err != nil {
		logger.Infof("ERROR delete state: %d, %s. %s", id, stateName, err)
		return err
	}
	return nil
}
