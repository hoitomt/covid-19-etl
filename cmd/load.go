package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

func Load(loadChan chan CovidCase, wg *sync.WaitGroup) {
	log.Println("Start Load")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		Config.DbUserName, Config.DbPassword, Config.DbName, Config.DbHost, Config.DbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database %s", err)
	}

	log.Printf("Connection String: %s", connStr)
	for covidCase := range loadChan {
		rows, err := db.Query("SELECT * from counties where fips = $1")
		if err != nil {
			log.Printf("Error retrieving rows %s", err)
		}

		var id int
		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				log.Println("bad row")
			}
			log.Println(id)
		}

		log.Printf("Load CovidCase %#v", covidCase.ToCaseSql())
	}
	log.Println("Data Load Complete")
	wg.Done()
}

// import (
// 	"database/sql"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	age := 21
// 	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

// }
