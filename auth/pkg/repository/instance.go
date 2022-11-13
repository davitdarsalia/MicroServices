package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func NewDatabaseInstance() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5433 user=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("\nFailed To Connect To The Database. \nOriginal Error: %s\n", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("\nPing Error. Connection Lost\n. \n Original Error: %s", err.Error())
	}
	return db, nil
}
