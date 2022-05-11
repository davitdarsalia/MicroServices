package service

import (
	"github.com/go-redis/redis/v8"

	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
)

// AuthService - AuthService needs access To DB. Initialize DB in the constructor
type AuthService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

func (s *AuthService) CheckUser(u *entities.UserInput) (int, error) {
	return 0, nil
}

/* Working with endpoint Methods */

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	hash, salt := generateHash(u.Password)

	// Implementing hashing and salting
	u.Password = hash
	u.Salt = salt

	//s.redisConn.Do("SET", "UniqueSalt", u.Salt)
	//reply, data := s.redisConn.Do("GET", "UniqueSalt")

	//fmt.Println(reply, data)

	// Moving user instance to the lower level - Repository level
	return s.repo.RegisterUser(u)
}
