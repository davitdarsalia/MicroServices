package repository

import "github.com/jmoiron/sqlx"

/* Database Related Types */

type Persistor interface {
	Authorizer
	Verifier
	Reset
}

type Repository struct {
	Persistor
}

type DBInstance struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{Persistor: &DBInstance{db: db}}
}

/* Root Interfaces */

type Authorizer interface {
	create()
	login()
	refresh()
}

type Verifier interface {
	verify()
}

type Reset interface {
	reset()
}
