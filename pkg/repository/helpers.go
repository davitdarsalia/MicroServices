package repository

import "github.com/jmoiron/sqlx"

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
