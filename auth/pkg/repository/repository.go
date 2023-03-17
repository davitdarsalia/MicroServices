package repository

import (
	"auth/internal/entities"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=AuthDB
type AuthDB interface {
	CreateUser(u entities.User) (string, error)
	LoginUser(u entities.UserInput) ([3]string, error)
	RecoverPassword(u entities.RecoverPasswordInput) error
}

type Repository struct {
	AuthDB
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func New(db *sqlx.DB) Repository {
	return Repository{AuthDB: NewAuthPostgres(db)}
}
