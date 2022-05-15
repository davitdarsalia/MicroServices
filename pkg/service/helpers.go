package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"
	"math/rand"
	"net"
	"os"
	"time"
)

var localContext = context.Background()

type AuthService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

type AccountService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

type TransactionsService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

type DeletionsService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

func NewAuthService(r repository.Authorization, redisConn *redis.Client) *AuthService {
	return &AuthService{repo: r, redisConn: redisConn}
}

func NewAccountService(r repository.Authorization, redisConn *redis.Client) *AccountService {
	return &AccountService{repo: r, redisConn: redisConn}
}
func NewTransactionsService(r repository.Authorization, redisConn *redis.Client) *AccountService {
	return &AccountService{repo: r, redisConn: redisConn}
}
func NewDeletionsService(r repository.Authorization, redisConn *redis.Client) *AccountService {
	return &AccountService{repo: r, redisConn: redisConn}
}

// Non Interface Methods

func generateHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateUniqueSalt(bytesAmount int) string {
	var saltBytes []byte

	for i := 0; i < 10; i++ {
		saltBytes = randstr.Bytes(bytesAmount)
	}
	return string(saltBytes)
}

func generateRandNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn((max - min + 1) + min)
}

func getIp() string {
	var result string

	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)
	for _, a := range address {
		if ipv4 := a.To4(); ipv4 != nil {
			result = fmt.Sprintf("IPv4: %s ", ipv4)
		}
	}

	return result
}
