package repository

import "github.com/jmoiron/sqlx"

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func NewAccountPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func NewTransactionsPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func NewDeletionsPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
