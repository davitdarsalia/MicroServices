package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

type claims struct {
	jwt.StandardClaims
	userId int
}

func JwtGenerator(userId string) string {
	var secretKey = []byte(os.Getenv("SIGN_KEY"))
	expiryDate, parseIntErr := strconv.ParseInt(os.Getenv("JWT_EXPIRY_DATE"), 0, 0)

	if parseIntErr != nil {
		log.Println(parseIntErr)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiryDate)).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Authentication Regular User",
	})

	token, err := claims.SignedString(secretKey)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

func TokenIsValid(tokenStr string) bool {
	key := os.Getenv("SIGN_KEY")
	token, tokenValidationErr := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if tokenValidationErr != nil {
		return false
	}

	return token.Valid
}
