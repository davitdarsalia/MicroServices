package service

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
)

func (s *AuthService) CheckUser(u *entities.UserInput) (int, error) {
	salt, redisGetError := s.redisConn.Get(localContext, "UniqueSalt").Result()

	if redisGetError != nil {
		log.Fatal(redisGetError)
	}

	fmt.Println(salt, "DDDD Salt")

	return 0, nil
}

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	hash, salt := generateHash(u.Password)

	u.Password = hash
	u.Salt = salt

	redisWriteErr := s.redisConn.Set(localContext, "UniqueSalt", salt, 0).Err()

	if redisWriteErr != nil {
		log.Fatal(redisWriteErr)
	}

	return s.repo.RegisterUser(u)
}
