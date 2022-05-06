package repository

import "github.com/davitdarsalia/LendAppBackend/pkg/service"

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

func NewRepository(service service.Service) *Repository {
	return &Repository{}
}
