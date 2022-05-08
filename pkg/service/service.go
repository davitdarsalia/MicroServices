package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
)

// Service - Contains all the provided services from our Bank Api
type Service struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
	// RegisterUser - Accepts User instance struct and returns its ID and an Error
	RegisterUser(u *entities.User) (int, error)
	// LoginUser - Accepts the same parameter and returns the same values as method below
	LoginUser(u entities.User) (int, error)
}

type Account interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Account:       nil,
		Transactions:  nil,
		Deletions:     nil,
	}
}
