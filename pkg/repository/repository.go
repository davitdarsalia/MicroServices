package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
}

type Account interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
