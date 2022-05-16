package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/dgrijalva/jwt-go"
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

func (s *AuthService) CheckUser(username, password string) (string, error) {
	salt, _ := s.redisConn.Get(localContext, "UniqueSalt").Result()
	user, err := s.repo.CheckUser(username, generateHash(password, salt))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.CustomToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			Id:        fmt.Sprintf("%d", user.UserID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   "Authentication",
		},
		UserID: user.UserID,
		Role:   "user",
		Ip:     getIp(),
	})

	return token.SignedString([]byte(entities.SignKey))
}

func (s *AuthService) ResetPassword(r *entities.ResetPassword) (string, error) {
	otp := generateResetEmail(r.Email)

	id, err := s.repo.ResetPassword(r)

	s.redisConn.Set(localContext, "OTP", otp, entities.OtpExpireDate)

	return id, err
}

func (s *AuthService) ValidateResetEmail() {

}

func (s *AuthService) RefreshLogin() {

}

func (s *AuthService) ResetPasswordProfile() {

}

func (s *AuthService) OtpGenerator() {

}
