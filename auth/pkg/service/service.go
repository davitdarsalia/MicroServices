package service

import (
	"auth/internal/entities"
	"auth/pkg/repository"
)

type Authorizer interface {
	CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error)
	LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error)
	RecoverPassword(u entities.RecoverPasswordInput) error
}

type Service struct {
	Authorizer
}

type AuthService struct {
	repo repository.AuthDB
}

func newAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func New(r repository.Repository) *Service {
	return &Service{Authorizer: newAuthService(r)}
}
