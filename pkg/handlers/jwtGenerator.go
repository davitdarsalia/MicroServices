package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
)

func JwtGenerator() string {
	var (
		signKey  = []byte(os.Getenv("SIGN_KEY"))
		audience = os.Getenv("JWT_AUDIENCE")
		id       = os.Getenv("JWT_ID")
	)

	expiryDate, parseIntErr := strconv.ParseInt(os.Getenv("JWT_EXPIRY_DATE"), 0, 0)

	if parseIntErr != nil {
		log.Println(parseIntErr)
	}

	claims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: expiryDate,
		Id:        id,
		IssuedAt:  1,
		Issuer:    "test",
		NotBefore: 0,
		Subject:   "Sign",
	}

	//
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, ssError := token.SignedString(signKey)

	if ssError != nil {
		log.Println(ssError)
	}

	return signedString
}
