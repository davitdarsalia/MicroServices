package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

func JwtGenerator() string {
	var (
		secret   = []byte(os.Getenv("SIGN_KEY"))
		audience = os.Getenv("JWT_AUDIENCE")
		issuer   = os.Getenv("ISSUER")
		subject  = os.Getenv("JWT_SUBJECT")
	)

	expiryDate, parseIntErr := strconv.ParseInt(os.Getenv("JWT_EXPIRY_DATE"), 0, 0)

	if parseIntErr != nil {
		log.Println(parseIntErr)
	}

	claims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiryDate)).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
		Subject:   subject,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString(secret)

	if err != nil {
		log.Fatal(err)
	}
	return strToken
}
