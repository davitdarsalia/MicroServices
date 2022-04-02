package db

import (
	"database/sql"
	"dbPractice/pkg/constants"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {
	db, dbOpenError := sql.Open("postgres", constants.DbConfig)
	if dbOpenError != nil {
		log.Fatal(dbOpenError)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
