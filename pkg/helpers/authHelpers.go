package helpers

import (
	"github.com/davitdarsalia/RestAPI.git/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

func JWTGenerator(userId string) models.TokenModel {
	var (
		secretKey = []byte(os.Getenv("SIGN_KEY"))
		expiresIn = os.Getenv("JWT_EXPIRY_DATE")
	)
	exDate, err := strconv.ParseInt(expiresIn, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(exDate)).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Regular User Authentication",
	})
	t, claimErr := jwtClaims.SignedString(secretKey)
	if claimErr != nil {
		log.Fatal(claimErr)
	}
	date, parseErr := strconv.ParseInt(expiresIn, 0, 0)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	token := models.TokenModel{
		AccessToken: t,
		ExpiresIn:   time.Duration(date),
		TokenType:   "Bearer",
		UserId:      userId,
	}
	return token
}
