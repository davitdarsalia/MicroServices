package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type claims struct {
	jwt.StandardClaims
	userId int
}

func JwtGenerator(w http.ResponseWriter, userId int) string {
	var secretKey = []byte(os.Getenv("SIGN_KEY"))
	intId := strconv.Itoa(userId)
	expiryDate, parseIntErr := strconv.ParseInt(os.Getenv("JWT_EXPIRY_DATE"), 0, 0)

	if parseIntErr != nil {
		log.Println(parseIntErr)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiryDate)).Unix(),
		Id:        "",
		IssuedAt:  time.Now().Unix(),
		Issuer:    intId,
		Subject:   "Authentication Regular User",
	})

	token, err := claims.SignedString(secretKey)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

// GetUserId - Helper Function, Which Returns User Id
func GetUserId(t string) int {
	key := os.Getenv("SIGN_KEY")

	tokenClaims := &claims{}
	_, parseErr := jwt.ParseWithClaims(t, tokenClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if parseErr != nil {
		log.Println(parseErr)
	}

	return tokenClaims.userId
}
