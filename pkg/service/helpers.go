package service

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/thanhpk/randstr"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var localContext = context.Background()
var localSendgridClient = sendgrid.NewSendClient(os.Getenv("MAIL_API_KEY"))

type AuthService struct {
	repo      repository.Authorization
	redisConn *redis.Client
}

type AccountService struct {
	repo      repository.Account
	redisConn *redis.Client
}

type SettingsService struct {
	repo      repository.Settings
	redisConn *redis.Client
}

func NewAuthService(r repository.Authorization, redisConn *redis.Client) *AuthService {
	return &AuthService{repo: r, redisConn: redisConn}
}

func NewAccountService(r repository.Account, redisConn *redis.Client) *AccountService {
	return &AccountService{repo: r, redisConn: redisConn}
}

func NewSettingsService(r repository.Settings, redisConn *redis.Client) *SettingsService {
	return &SettingsService{repo: r, redisConn: redisConn}
}

// Non Interface Methods

func (s *AuthService) GenerateToken(userID int) (string, error) {
	e, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.CustomToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(e)).Unix(),
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

func (s *AuthService) GenerateSessionID() string {
	userID, _ := s.redisConn.Get(localContext, constants.RedisID).Result()
	i := sha1.New()
	i.Write([]byte(userID))

	return fmt.Sprintf("%x", i.Sum([]byte(generateUniqueSalt(20))))

}

func generateUniqueSalt(bytesAmount int) string {
	var saltBytes []byte

	for i := 0; i < 10; i++ {
		saltBytes = randstr.Bytes(bytesAmount)
	}
	return string(saltBytes)
}

func generateHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func generateRandNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn((max - min + 1) + min)
}

func generateEmail(sendTo *string, message, subject, plainText string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER"), os.Getenv("MAIL_FROM"))

	to := mail.NewEmail("Receiver", *sendTo)

	html := fmt.Sprintf(
		"<strong>%s</strong>",
		message,
	)

	mail := mail.NewSingleEmail(from, subject, to, plainText, html)

	res, emailWriterErr := localSendgridClient.Send(mail)

	if emailWriterErr != nil {
		log.Printf("Email Helper Method: %s", emailWriterErr.Error())
	}

	log.Printf("Email Logger: %d, %s, %s", res.StatusCode, res.Body, res.Headers)

}

func generateResetEmail(sendTo string) string {
	otp := generateOTP()

	// SendGrid Stuff
	from := mail.NewEmail("Bene Store", os.Getenv("MAIL_FROM"))
	to := mail.NewEmail("Receiver", sendTo)

	subject := "Bene Store Reset Password"
	text := "Verification Code:"

	html := fmt.Sprintf("<strong>%s</strong>", otp)

	mail := mail.NewSingleEmail(from, subject, to, text, html)

	res, emailWriterErr := localSendgridClient.Send(mail)

	if emailWriterErr != nil {
		log.Printf("Email Helper Method: %s", emailWriterErr.Error())
	}

	log.Printf("Email Logger: %d, %s, %s", res.StatusCode, res.Body, res.Headers)

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

func formatNowDate() string {
	return time.Now().Format(entities.RegularFormat)
}

func (a *AccountService) getRedisUserID() int {
	id, err := a.redisConn.Get(localContext, constants.RedisID).Result()

	if err != nil {
		log.Printf("%s : %s", err, "RedisGetError")
	}

	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("%s : %s", err, "[Redis] - ParseInt Error")
	}

	return intID

}

func (s *SettingsService) getRedisID() int {
	id, err := s.redisConn.Get(localContext, constants.RedisID).Result()

	if err != nil {
		log.Printf("%s : %s", err, "RedisGetError")
	}

	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("%s : %s", err, "[Redis] - ParseInt Error")
	}

	return intID
}

func validateRegFields(u *entities.User) bool {
	checks := []string{emailRegex, usernameRegex, nameRegex, passwordRegex, ipAddressRegex, mobileNumberRegex, personalNumberRegex}
	values := []string{u.Email, u.UserName, fmt.Sprintf("%s %s", u.FirstName, u.LastName), u.Password, u.IpAddress, u.PhoneNumber, u.PersonalNumber}

	var res bool

	for i, v := range checks {
		r, err := regexp.Compile(v)

		if err != nil {
			log.Printf("Regex Compile Error: %s", err.Error())
		}

		match := r.MatchString(values[i])

		if match == false {
			res = false
			break
		} else {
			res = true
		}
	}

	return res
}
