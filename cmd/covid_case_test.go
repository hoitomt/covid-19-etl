package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCovidCountyCase(t *testing.T) {
	testData := []string{"2020-01-21", "Snohomish", "Washington", "53061", "1", "0"}
	countyCase := NewCountyCase(testData)

	assert.NotEqual(t, nil, *countyCase.Date)
	assert.Equal(t, "2020-01-21", countyCase.Date.Format("2006-01-02"))
	assert.Equal(t, "Snohomish", countyCase.County)
	countyCase.State = "Washington"
	countyCase.Fips = "53061"
	countyCase.Cases = 1
	countyCase.Deaths = 0
}

func TestCovidStateCase(t *testing.T) {
	testData := []string{"2020-01-22", "Washington", "53", "11", "0"}
	stateCase := NewStateCase(testData)

	assert.NotEqual(t, nil, *stateCase.Date)
	assert.Equal(t, "2020-01-22", stateCase.Date.Format("2006-01-02"))
	stateCase.State = "Washington"
	stateCase.Fips = "53"
	stateCase.Cases = 11
	stateCase.Deaths = 0
}
