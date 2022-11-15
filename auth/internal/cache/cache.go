package cache

import "github.com/go-redis/redis/v8"

func New(c *redis.Options) *redis.Client {
	return redis.NewClient(c)
}
