package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	var err error

	var connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	for i := 0; i <= 5; i++ {

		DB, err = sql.Open("postgres", connStr)

		if err == nil {
			err = DB.Ping()

			if err == nil {
				log.Println("Connected to DB")
				return
			}
		}

		log.Println("Retrying to connect DB")

		time.Sleep(2 * time.Second)

	}
	log.Fatal("Could not connect", err)
}
