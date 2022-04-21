package helpers

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func DbConnection() *sqlx.DB {
	c := os.Getenv("DB_CONFIG")
	db, err := sqlx.Open("postgres", c)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
