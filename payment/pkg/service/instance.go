package service

import (
	"github.com/davitdarsalia/payment/pkg/repository"
	"github.com/go-redis/redis/v8"
)

func NewServiceInstance(r repository.Repository, redis *redis.Client) *RootService {
	return &RootService{repository: r, redis: redis}
}
