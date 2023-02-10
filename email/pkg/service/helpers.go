package service

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "golang.org/x/crypto/bcrypt"
	"math/rand"
	"net"
	"os"
	"time"
)

// TODO - move token expiry time in .env or config
const expiry = 200

func accessToken(userId, signKey string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiry * time.Minute).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Authorization",
	})

	return t.SignedString([]byte(signKey))
}

func refreshToken(userId, signKey string) (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiry * time.Hour).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Authentication ",
	})

	return t.SignedString([]byte(signKey))
}

func hash(password, salt string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	b := hasher.Sum([]byte(salt))
	return hex.EncodeToString(b)
}

func getIPv6() (result string) {
	host, _ := os.Hostname()
	address, _ := net.LookupIP(host)

	for _, a := range address {
		if ipv16 := a.To16(); ipv16 != nil {
			result = fmt.Sprintf("%s ", ipv16)
		}
	}

	return result
}

func getFormattedDateTime() string {
	return time.Now().Format("2006-01-02 15:04")
}

func generateSalt() (string, error) {
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
