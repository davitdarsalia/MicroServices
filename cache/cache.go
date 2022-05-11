package cache

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisCache(c *redis.Options) *redis.Client {
	client := redis.NewClient(c)

	return client
}
