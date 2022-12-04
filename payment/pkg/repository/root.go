package repository

import (
	"github.com/jackc/pgx/v5"
)

/* Database Related Types */

type Payments interface {
	TestPayments
	RealPayments
}

type Repository struct {
	Payments
}

type DBInstance struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{Payments: &DBInstance{db: db}}
}

/* Root Interfaces */

type TestPayments interface {
	FetchPublicKey() string
}

type RealPayments interface {
	FetchPublicKey() string
}
