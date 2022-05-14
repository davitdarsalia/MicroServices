package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
)

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	amountOfBytes := generateRandNumber(5, 20)
	salt := generateUniqueSalt(amountOfBytes)
	hash := generateHash(u.Password, salt)

	u.Password = hash
	u.Salt = []byte(salt)

	redisWriteErr := s.redisConn.Set(localContext, "UniqueSalt", salt, 0).Err()

	if redisWriteErr != nil {
		log.Fatal(redisWriteErr)
	}

	return s.repo.RegisterUser(u)
}

func (s *AuthService) CheckUser(u *entities.UserInput) (int, error) {
	salt, redisGetError := s.redisConn.Get(localContext, "UniqueSalt").Result()

	if redisGetError != nil {
		log.Fatal(redisGetError)
	}

	hash := generateHash(u.Password, salt)

	u.Password = hash

	return s.repo.CheckUser(u.UserName, u.Password)
}
