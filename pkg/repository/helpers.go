package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Account
	Settings
}

type AuthPostgres struct {
	db *sqlx.DB
}

type AccountPostgres struct {
	db *sqlx.DB
}

type SettingsPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func NewSettingsPostgres(db *sqlx.DB) *SettingsPostgres {
	return &SettingsPostgres{db: db}
}
