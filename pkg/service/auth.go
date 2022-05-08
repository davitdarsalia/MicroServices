package service

import (
	"crypto/sha512"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/thanhpk/randstr"
)

// AuthService - AuthService needs access To DB. Initialize DB in the constructor
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{repo: r}
}

/* Working with endpoint Methods */

func (s *AuthService) LoginUser(u entities.User) (int, error) {
	//TODO implement me
	panic(any("DDD"))
}

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	hash, salt := s.generateHash(u.Password)

	// Implementing hashing and salting
	u.Password = hash
	u.Salt = salt

	// Moving user instance to the lower level - Repository level
	return s.repo.RegisterUser(u)
}

/* Helper methods */

// generateHash - Returns an actual hash string + slice of bytes
// which will be stored in DataBase as a unique salt
func (s *AuthService) generateHash(password string) (string, []byte) {
	// Define nullish constructor
	hash := sha512.New()

	salt := generateUniqueSalt(255)
	// Write password with slice of bytes
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
