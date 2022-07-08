package service

import (
	"errors"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
	"strconv"
)

func (s *AuthService) RegisterUser(u *entities.User) (int, error) {
	amountOfBytes := generateRandNumber(5, 20)
	salt := generateUniqueSalt(amountOfBytes)
	hash := generateHash(u.Password, salt)

	u.Password = hash
	u.Salt = []byte(salt)

	userID, err := s.repo.RegisterUser(u)

	if err == nil {
		// Auth Bugfix - If User IS Already Registered Salt Will Not Stored In Redis. This Prevents 404 Error At CheckUser Method
		redisWriteErr := s.redisConn.Set(localContext, "UniqueSalt", salt, 0).Err()
		if redisWriteErr != nil {
			log.Printf("RedisWriterError: %v", err.Error())
		}
	}

	return userID, err

}

func (s *AuthService) CheckUser(username, password string) (string, error) {
	c := make(chan string)
	defer func() {
		close(c)
	}()

	go func() {
		salt, _ := s.redisConn.Get(localContext, "UniqueSalt").Result()
		c <- salt
	}()

	user, err := s.repo.CheckUser(username, generateHash(password, <-c))

	// TODO - Clear all previous redis keys

	s.redisConn.Set(localContext, constants.RedisID, user.UserID, 0)

	// TODO - Ensure, that redis key can be appendable only one time with InvalidID constant

	if err != nil {
		return "", err
	}

	if err == nil {
		go s.redisConn.Set(localContext, constants.SessionID, s.GenerateSessionID(), 0)
	}

	return entities.GenerateToken(user.UserID)
}

func (s *AuthService) ResetPassword(r *entities.ResetPassword) (string, error) {
	otp := generateResetEmail(r.Email)

	id, err := s.repo.ResetPassword(r)

	s.redisConn.Set(localContext, "OTP", otp, entities.OtpExpireDate)

	return id, err
}

func (s *AuthService) ValidateResetEmail(e *entities.ValidateResetEmail) error {
	otp, _ := s.redisConn.Get(localContext, "OTP").Result()
	salt, _ := s.redisConn.Get(localContext, "UniqueSalt").Result()

	e.NewPassword = generateHash(e.NewPassword, salt)

	if e.ValidationCode != otp {
		return errors.New("incorrect OTP Code")
	}

	err := s.repo.ValidateResetEmail(e)
	return err

}

func (s *AuthService) ResetPasswordProfile(e *entities.ResetPasswordInput) error {
	c := make(chan string, 1)
	defer func() {
		close(c)
	}()

	go func() {
		salt, _ := s.redisConn.Get(localContext, "UniqueSalt").Result()
		c <- salt
	}()

	e.NewPassword = generateHash(e.NewPassword, <-c)

	err := s.repo.ResetPasswordProfile(e)

	return err

}

func (s *AuthService) RefreshLogin() int {
	c := make(chan string, 1)
	defer func() {
		close(c)
	}()

	go func() {
		id, err := s.redisConn.Get(localContext, constants.RedisID).Result()

		c <- id

		if err != nil {
			log.Fatal("Redis Get UserID Error")
		}

	}()

	intID, err := strconv.Atoi(<-c)

	if err != nil {
		log.Fatal("ParseInt Error")
	}

	return intID

}

func (s *AuthService) OtpGenerator() {

}
