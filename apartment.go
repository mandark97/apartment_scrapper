package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Apartment is nice
type Apartment struct {
	Title              string
	DateAdded          time.Time
	OfferedBy          string
	Surface            int
	Partitioning       string
	YearOfConstruction string
	Floor              string
	NoRooms            int
	Description        string
	Images             []string
	Price              string
	Link               string
}

func dbWrite(apartment Apartment) {
	const (
		host = "localhost"
		port = 5432
		user = "mandark"
		// password = ""
		dbname = "lucas_db"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// some debug print logs
	log.Print("Successfully connected!")
	// fmt.Printf("%s, %s, %s, %f", product.Name, product.Code, product.Description, product.Price)
	// sqlStatement := `
	// INSERT INTO floryday (product, code, description, price)
	// VALUES ($1, $2, $3, $4)`
	// _, err = db.Exec(sqlStatement, product.Name, product.Code, product.Description, product.Price)
	// if err != nil {
	// 	panic(err)
	// }
}
