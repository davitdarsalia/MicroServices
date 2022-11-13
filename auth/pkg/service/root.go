package service

import "github.com/davitdarsalia/auth/pkg/repository"

type ProviderService interface {
	Create()
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
