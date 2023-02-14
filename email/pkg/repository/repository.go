package repository

import (
	"github.com/jmoiron/sqlx"
)

type EmailDB interface {
	Dummy()
}

type Repository struct {
	EmailDB
}

type EmailPostgres struct {
	db *sqlx.DB
}

func NewEmailDB(db *sqlx.DB) *EmailPostgres {
	return &EmailPostgres{db: db}
}

func New(db *sqlx.DB) Repository {
	return Repository{EmailDB: NewEmailDB(db)}
}
