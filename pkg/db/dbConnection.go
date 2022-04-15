package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	dbConfig := os.Getenv("DB_CONFIG")
	db, dbOpenError := sql.Open("postgres", dbConfig)
	if dbOpenError != nil {
		log.Fatal(dbOpenError)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
