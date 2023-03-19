package repository

import (
	"auth/internal/entities"
	"github.com/jackc/pgx/v4/pgxpool"
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
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func New(db *pgxpool.Pool) Repository {
	return Repository{AuthDB: NewAuthPostgres(db)}
}
