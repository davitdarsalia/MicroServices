package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"menuAPI/internal/entities"
)

func NewDB(c entities.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		// TODO -  Insert password field
		"postgres", fmt.Sprintf("host=%s port=%s user=%s sslmode=%s",
			c.Host, c.Port, c.Username, c.SSLMode),
	)

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
