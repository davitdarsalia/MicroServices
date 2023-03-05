package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", os.Getenv("DB_CONN_STRING"))

	// TODO - Execute migrations

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
