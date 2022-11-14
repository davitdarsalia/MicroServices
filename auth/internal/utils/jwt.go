package utils

import (
	"fmt"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

// TokenPair - Returns Token Pair (Access Token + Refresh Token) And Corresponding Error
func TokenPair(userID types.UserID) ([2]string, error) {
	tokenPairChan := make(chan string, 2)

	m, _ := strconv.Atoi(os.Getenv("EXP"))
	exp := time.Minute * time.Duration(m)

	aT := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.AccessToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
			Id:        fmt.Sprintf("TokenID: %s", Salt()),
			IssuedAt:  time.Now().Unix(),
			Subject:   "Authentication",
		},
		UserID: userID,
		Role:   "User",
		Ip:     IpAddress(),
	})

	rT := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7200).Unix(),
			Id:        fmt.Sprintf("TokenID: %s", Salt()),
			IssuedAt:  time.Now().Unix(),
			Subject:   "Refresh"},
	)

	go func() {
		accessToken, err := aT.SignedString([]byte(os.Getenv("SIGN_KEY")))

		if err != nil {
			log.Fatal(err)
		}

		tokenPairChan <- accessToken
	}()

	go func() {
		refreshToken, err := rT.SignedString([]byte(os.Getenv("SIGN_KEY")))

		if err != nil {
			log.Fatal(err)
		}

		tokenPairChan <- refreshToken
	}()

	return [2]string{<-tokenPairChan, <-tokenPairChan}, nil
}
