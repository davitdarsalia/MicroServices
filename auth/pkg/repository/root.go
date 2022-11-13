package repository

import (
	"github.com/jackc/pgx/v5"
)

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
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{Persistor: &DBInstance{db: db}}
}

/* Root Interfaces */

type Authorizer interface {
	Create()
	Login()
	Refresh()
}

type Verifier interface {
	Verify()
}

type Reset interface {
	Reset()
}
