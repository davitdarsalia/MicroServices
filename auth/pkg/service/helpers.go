package service

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "golang.org/x/crypto/bcrypt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

// TODO - move token expiry time in .env or config
const expiry = 200
const notAuthorizedResponse = "Not authorized"

func generateValidationStruct(e error) error {
	fieldNames := make([]string, 0, 0)
	re := regexp.MustCompile(`'([^']*)' failed`)
	matches := re.FindAllStringSubmatch(e.Error(), -1)

	for _, match := range matches {
		fieldNames = append(fieldNames, match[1])
	}

	return errors.New(fmt.Sprintf("verifications failed for fields: %v", fieldNames))
}

func checkToken(authToken, signKey string) (string, error) {
	tok, err := jwt.ParseWithClaims(authToken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return notAuthorizedResponse, err
	}

	claims, ok := tok.Claims.(*jwt.StandardClaims)

	if !ok {
		return notAuthorizedResponse, errors.New("invalid token claims")
	}

	return claims.Id, nil

}

func accessToken(userId, signKey string) (string, error) {
	var length = len(strings.Split(signKey, ""))

	if length < 20 || length > 40 {
		return "", errors.New("invalid sign key")
	}

	if ok := checkUUID(userId); !ok {
		return "", errors.New("invalid user id. User id is not correct uuid format")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiry * time.Minute).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Subject:   "Authorization",
	})

	return t.SignedString([]byte(signKey))
}

func checkUUID(uuid string) bool {
	condition := regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	return condition.MatchString(uuid)
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

	// hasher.Size()

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
