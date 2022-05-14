package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/thanhpk/randstr"
	"math/rand"
	"net"
	"os"
	"time"
)

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.CustomToken{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			Id:        fmt.Sprintf("%d", userID),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("ISSUER"),
			Subject:   "Authentication",
		},
		Role: "user",
		Ip:   getIp(),
	})

	byteSignature := randstr.Bytes(rand.Intn(55))
	return token.SignedString(byteSignature)
}

func getIp() string {
	var result string

	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)
	for _, a := range address {
		if ipv4 := a.To4(); ipv4 != nil {
			result = fmt.Sprintf("IPv4: %s ", ipv4)
		}
	}

	return result
}
