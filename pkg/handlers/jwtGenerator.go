package handlers

import (
	"dbPractice/pkg/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

func JwtGenerator(userId string) models.TokenModel {
	var secretKey = []byte(os.Getenv("SIGN_KEY"))
	var expiresIn = os.Getenv("JWT_EXPIRY_DATE")
	expiryDate, parseIntErr := strconv.ParseInt(expiresIn, 0, 0)
	if parseIntErr != nil {
		log.Println(parseIntErr)
	}
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiryDate)).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Authentication Regular User",
	})
	token, err := jwtClaims.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	date, parseErr := strconv.ParseInt(expiresIn, 0, 0)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	var t = models.TokenModel{
		AccessToken: token,
		ExpiresIn:   date,
		TokenType:   "Bearer",
		UserId:      userId,
	}

	return t
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

// Refactor This

func RefreshLogin(token string, userId string) (string, error) {
	valid := TokenIsValid(token)
	fmt.Println(valid)

	//if valid {
	//	newToken := JwtGenerator(userId)
	//	return newToken, nil
	//}
	//newToken := JwtGenerator(userId)
	//return newToken, nil
	return "", nil
}
