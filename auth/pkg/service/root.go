package service

import (
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/pkg/repository"
)

type ProviderService interface {
	Create(u *entities.User) (string, error)
	Login()
	Refresh()
	Verify()
	Reset()
}

type Service struct {
	ProviderService
}

type RootService struct {
	repository repository.Repository
}

func New(repos *repository.Repository) *Service {
	return &Service{ProviderService: NewServiceInstance(*repos)}
}
