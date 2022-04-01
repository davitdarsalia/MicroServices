package db

import (
	"database/sql"
	"dbPractice/pkg/constants"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func DbConnection() *sql.DB {
	defer fmt.Println("Db Connected")
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
