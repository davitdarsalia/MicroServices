package service

import (
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/davitdarsalia/auth/pkg/repository"
	"github.com/go-redis/redis/v8"
)

type ProviderService interface {
	Create(u *entities.User) (string, error)
	Login(u *entities.UserInput) (types.TokenPair, error)
	Refresh(refreshToken types.RefreshToken) (types.TokenPair, error)
	Reset(u *entities.ResetPasswordInput) error
	Verify()
}

type Service struct {
	ProviderService
}

type RootService struct {
	repository repository.Repository
	redis      *redis.Client
}

func New(repos *repository.Repository, redis *redis.Client) *Service {
	return &Service{ProviderService: NewServiceInstance(*repos, redis)}
}
