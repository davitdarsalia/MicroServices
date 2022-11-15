package utils

import (
	"errors"
	"fmt"
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

// TokenPair - Returns Token Pair (Access Token + Refresh Token) And Corresponding Error
func TokenPair(userID types.UserID) types.TokenPair {
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
		entities.RefreshToken{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(7200).Unix(),
				Id:        fmt.Sprintf("TokenID: %s", Salt()),
				IssuedAt:  time.Now().Unix(),
				Subject:   "Refresh",
			},
			Ip: IpAddress(),
		},
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

	return [2]string{<-tokenPairChan, <-tokenPairChan}
}

func ParseAccessToken(t string) (int, error) {
	token, err := jwt.ParseWithClaims(t, &entities.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		// TODO - Replace Static Sign Key With Custom One For Each User
		return []byte(os.Getenv("sign_key")), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*entities.RefreshToken)
	if !ok {
		return -1, errors.New(constants.InvalidTokenClaims)
	}
	userID, _ := strconv.Atoi(claims.Id)
	return userID, nil
}

func ParseRefreshToken(t string) (types.TokenPair, error) {
	token, err := jwt.ParseWithClaims(t, &entities.RefreshToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Signing Method")
		}
		// TODO - Replace Static Sign Key With Custom One For Each User
		return []byte(os.Getenv("sign_key")), nil
	})
	if err != nil {
		return [2]string{"", ""}, err
	}
	claims, ok := token.Claims.(*entities.RefreshToken)
	if !ok {
		return [2]string{"", ""}, errors.New(constants.InvalidTokenClaims)
	}

	return TokenPair(claims.Id), nil
}
