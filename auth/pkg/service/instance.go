package service

import "github.com/davitdarsalia/auth/pkg/repository"

func NewServiceInstance(r repository.Repository) *RootService {
	return &RootService{repository: r}
}
