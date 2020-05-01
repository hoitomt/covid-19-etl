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

func (db *CovidDb) GetCountyByName(countyName string) (*County, error) {
	var err error

	county := County{}
	err = db.Get(&county, "SELECT * FROM counties WHERE lower(name) = $1", strings.ToLower(countyName))

	switch err {
	case sql.ErrNoRows:
		logger.Infof("No rows returned from County query.")
		return nil, nil
	case nil:
		return &county, nil
	default:
		logger.Fatalf("Bad things man - Select Id2 %s", err)
		return nil, err
	}
}

func (db *CovidDb) InsertCounty(county County) error {
	var id int
	rows, err := db.NamedQuery("INSERT INTO counties (name, fips, state_id) VALUES (:name, :fips, :state_id) RETURNING id", county)
	if err != nil {
		logger.Errorf("ERROR inserting county: %v. %s", county, err)
		return err
	}
	if rows.Next() {
		rows.Scan(&id)
	}
	county.Id = id
	return nil
}

func (db *CovidDb) DeleteCounty(id int, countyName string) error {
	var err error
	if id > 0 {
		_, err = db.NamedQuery("DELETE FROM counties where id = :id", id)
	} else {
		_, err = db.NamedQuery("DELETE FROM counties where lower(name) = :name", countyName)
	}
	if err != nil {
		logger.Infof("ERROR delete county: %d, %s. %s", id, countyName, err)
		return err
	}
	return nil
}
