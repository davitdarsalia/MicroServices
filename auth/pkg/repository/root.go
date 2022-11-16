package repository

import (
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
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
	Create(u *entities.User) (string, error)
	Login(email, password string) (string, error)
}

type Verifier interface {
	Verify()
}

type Reset interface {
	Reset(email, idNumber string, newPassword types.Hash512) error
}
