package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"
	"log"
	"math/rand"
	"net/smtp"
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

func (s *AuthService) GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.CustomToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			Id:        fmt.Sprintf("%d", userID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   "Authentication",
		},
		UserID: userID,
		Role:   "user",
		Ip:     entities.GetIp(),
	})
	return token.SignedString([]byte(entities.SignKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &entities.CustomToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		return []byte(entities.SignKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(*entities.CustomToken)

	if !ok {
		return 0, errors.New("invalid Token Claims")
	}

	return claims.UserID, nil
}

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

func generateResetEmail(sendTo ...string) string {
	otp := generateOTP()

	address := fmt.Sprintf("%s:%s", entities.MailHost, entities.MailPort)
	auth := smtp.PlainAuth("", entities.SendMailFrom, entities.MailAuthPassword, entities.MailHost)
	err := smtp.SendMail(address, auth, entities.SendMailFrom, sendTo, []byte(otp))

	if err != nil {
		log.Fatal(err)
	}

	return otp
}

func generateOTP() (otp string) {
	const (
		min = 100000
		max = 999999
	)

	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(max - min + 1)
	otp = fmt.Sprintf("%d", s)

	return
}
