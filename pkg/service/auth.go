package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	todo "github.com/davitdarsalia/BookStoreMicroservices"
	"github.com/davitdarsalia/BookStoreMicroservices/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	salt       = "r fjn rj21ned31ngb4j1en d1jd 4cj t54jg 2je"
	signInKey  = "bd8a063ed6cff56ceff7ca31239b905bedf42e9f70d0b8dca23dbbe8baca9968"
	expiryTime = 20 * time.Minute
)

type AuthService struct {
	repo repository.Authorization
}

type tokenCustomClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

/* ALl Methods Starts Here */

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiryTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
		Role:   "Customer",
	})

	// Returns complete signed string. Accepts as an argument salt key
	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// 1. Checking SignMethod
	token, err := jwt.ParseWithClaims(accessToken, &tokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Sign Method")
		}
		// This Method Returns A SignInMethod
		return []byte(signInKey), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// 2. Checking For Proper Claims
	claims, ok := token.Claims.(*tokenCustomClaims)

	if !ok {
		return 0, errors.New("token claims are not of type")
	}

	return claims.UserId, nil
}

// generatePasswordHash - Helper Function Based On Sha1
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
