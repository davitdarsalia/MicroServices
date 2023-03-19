package service

import (
	"auth/internal/entities"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
	_ "golang.org/x/crypto/bcrypt"
	"net"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"
)

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

func accessToken(pKey []byte, userID string) (string, error) {
	m, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRY_TIME"))
	if err != nil {
		return "", err
	}

	claims := entities.AccessToken{
		StandardClaims: jwt.StandardClaims{
			Audience:  "Regular User",
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(m)).Unix(),
			Issuer:    "Auth Server",
			NotBefore: 0,
			Subject:   "Authorization, Authentication",
		},
		AccessTokenCustomClaims: entities.AccessTokenCustomClaims{
			UserID:         userID,
			CreatedAt:      time.Now().String(),
			UserRole:       "User",
			ExpirationTime: fmt.Sprintf("%d Minutes", m),
			IpAddress:      getIPv6(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(pKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// refreshToken - TODO - Add signature
func refreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}
	token := base64.RawURLEncoding.EncodeToString(tokenBytes)
	return token, nil
}

func checkUUID(uuid string) bool {
	condition := regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	return condition.MatchString(uuid)
}

func hash(password, salt string) (string, error) {
	const timeCost = 1
	const memoryCost = 64 * 1024
	const threads = 4
	const keyLen = 32

	hash := argon2.IDKey([]byte(password), []byte(salt), timeCost, memoryCost, threads, keyLen)
	if hash == nil {
		return "", errors.New("failed to generate hash")
	}

	hashRunes := []rune(hex.EncodeToString(hash))

	// Use a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	for i := 0; i < len(hashRunes); i++ {
		wg.Add(1)
		go func(i int) {
			// Replace any invalid UTF-8 characters with a placeholder
			if !utf8.ValidRune(hashRunes[i]) {
				hashRunes[i] = '�'
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Convert the slice of runes back to a string
	hashString := string(hashRunes)
	return hashString, nil

}

func salt() ([]byte, error) {
	const length = 30
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[];:<>,.?/~`"

	var saltBytes = make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return nil, err
	}

	for i, b := range saltBytes {
		saltBytes[i] = charset[int(b)%len(charset)]
	}

	return saltBytes, nil
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
