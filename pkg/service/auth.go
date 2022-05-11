package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/gomodule/redigo/redis"
	"github.com/thanhpk/randstr"
)

// AuthService - AuthService needs access To DB. Initialize DB in the constructor
type AuthService struct {
	repo      repository.Authorization
	redisConn redis.Conn
}

func (s *AuthService) LoginUser(u *entities.UserInput) (int, error) {
	fmt.Println(u)

	s.redisConn.Do("get", "UniqueSalt")

	return 0, nil
}

func NewAuthService(r repository.Authorization, redisConn redis.Conn) *AuthService {
	return &AuthService{repo: r, redisConn: redisConn}
}

/* Working with endpoint Methods */

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	hash, salt := s.generateHash(u.Password)

	// Implementing hashing and salting
	u.Password = hash
	u.Salt = salt

	s.redisConn.Do("SET", "UniqueSalt", u.Salt)
	reply, data := s.redisConn.Do("GET", "UniqueSalt")

	fmt.Println(reply, data)

	// Moving user instance to the lower level - Repository level
	return s.repo.RegisterUser(u)
}

/* Helper methods */

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//p, _ := s.generateHash(password)
	////user, err := s.repo.CheckUser(username, p)
	//
	////if err != nil {
	////	return "", err
	////}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Minute * 3).Unix(),
	//	IssuedAt:  time.Now().Unix(),
	//})
	//
	//// Signed string argument must be unique for all users
	//return token.SignedString(generateUniqueSalt(15))
	return "", nil
}

// generateHash - Returns an actual hash string + slice of bytes
// which will be stored in DataBase as a unique salt
func (s *AuthService) generateHash(password string) (string, []byte) {
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
