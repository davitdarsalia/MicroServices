package service

import (
	"crypto/sha256"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"
)

// NewAuthService - Root Auth instance creator for Root Service Interface
func NewAuthService(r repository.Authorization, redisConn *redis.Client) *AuthService {
	return &AuthService{repo: r, redisConn: redisConn}
}

func GenerateToken(username, password string) (string, error) {
	return "", nil
}

func generateHash(password string) (string, []byte) {
	hash := sha256.New()
	salt := generateUniqueSalt(20)
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(salt)), salt
}

func generateUniqueSalt(bytesAmount int) []byte {
	var saltBytes []byte

	for i := 0; i < 10; i++ {
		saltBytes = randstr.Bytes(bytesAmount)
	}
	return saltBytes
}
