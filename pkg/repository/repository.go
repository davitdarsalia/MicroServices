package repository

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(username, password string) (entities.User, error)
	refreshLogin()
	resetPassword()
	resetPasswordProfile()
	otpGenerator()
}

type Account interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountPostgres(db),
		Transactions:  NewTransactionsPostgres(db),
		Deletions:     NewDeletionsPostgres(db),
	}
}
