package repository

import (
	"auth/internal/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

//go:generate mockery --name=AuthDB
type AuthDB interface {
	CreateUser(u *entities.User) (string, error)
	LoginUser(u entities.UserInput) (entities.UserMetaInfo, error)
	RequestPasswordRecover(u *entities.RecoverPasswordInput) (string, error)
	ResetPassword(u *entities.RecoverPasswordInput) (string, error)
}

type Repository struct {
	AuthDB
}

type AuthPostgres struct {
	db          *pgxpool.Pool
	credentials *entities.AWSCredentials
}

func NewAuthPostgres(db *pgxpool.Pool, c *entities.AWSCredentials) *AuthPostgres {
	return &AuthPostgres{db: db, credentials: c}
}

func New(db *pgxpool.Pool, c *entities.AWSCredentials) Repository {
	return Repository{AuthDB: NewAuthPostgres(db, c)}
}
