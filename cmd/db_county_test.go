package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func PrintMyPath() {
	fmt.Println(basepath)
}

func TestCovidCountyCase(t *testing.T) {
	initConfig()
	initLogger()

	db := NewDB()

	logger.Infof("DB: %v", db)

	county := db.GetCountyByName("eau claire")
	assert.Nil(t, county)

	// testData := []string{"2020-01-21", "Snohomish", "Washington", "53061", "1", "0"}
	// countyCase := NewCountyCase(testData)

	// assert.NotEqual(t, nil, *countyCase.Date)
	// assert.Equal(t, "2020-01-21", countyCase.Date.Format("2006-01-02"))
	// assert.Equal(t, "Snohomish", countyCase.County)
	// countyCase.State = "Washington"
	// countyCase.Fips = "53061"
	// countyCase.Cases = 1
	// countyCase.Deaths = 0
}
