package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(u *entities.UserInput) (int, error)
}

type Account interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewService(repos *repository.Repository, redisConn *redis.Client) *Service {
	return &Service{
		Authorization: NewAuthService(repos, redisConn),
		Account:       nil,
		Transactions:  nil,
		Deletions:     nil,
	}
}
