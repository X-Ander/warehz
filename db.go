package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func prepareDB() {
	var err error
	if db, err = sql.Open("mysql", dsn); err != nil {
		log.Fatal(err)
	}
}
