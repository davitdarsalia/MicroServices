package service

import "github.com/davitdarsalia/LendAppBackend/pkg/repository"

// Service - Contains all the provided services from our Bank Api
type Service struct {
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

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
