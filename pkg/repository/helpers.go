package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	Account
	Settings
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func NewAccountPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func NewSettingsPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
