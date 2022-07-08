package entities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net"
	"os"
	"strconv"
	"time"
)

func GenerateToken(userID int) (string, error) {
	exp := os.Getenv("ACCESS_TOKEN_EXP")
	intExp, _ := strconv.Atoi(exp)

	duration := time.Minute * time.Duration(intExp)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Id:        fmt.Sprintf("%d", userID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   "Authentication",
		},
		UserID: userID,
		Role:   "user",
		Ip:     GetIp(),
	})
	return token.SignedString([]byte(SignKey))
}

func GetIp() string {
	var result string

	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)
	for _, a := range address {
		if ipv4 := a.To4(); ipv4 != nil {
			result = fmt.Sprintf("%s ", ipv4)
		}
	}

	return result
}
