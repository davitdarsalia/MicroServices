package service

import (
	"github.com/davitdarsalia/payment/pkg/repository"
	"github.com/go-redis/redis/v8"
)

type ProviderService interface {
	FetchPublicKey() string
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
