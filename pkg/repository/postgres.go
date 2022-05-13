package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres  sslmode=disable")
	if err != nil {
		log.Fatalf("Failed To Connect DB. %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Ping Error. Connection Lost. %s", err.Error())
	}
	return db, nil
}
