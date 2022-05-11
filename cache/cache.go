package cache

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func NewRedisCache() redis.Conn {
	var (
		connType  = "tcp"
		redisPort = "localhost:6379"
	)

	redisInstance, err := redis.Dial(connType, redisPort)

	if err != nil {
		log.Fatalf("Error Initializing Redis Connection: %s", err.Error())
	}

	return redisInstance
}
