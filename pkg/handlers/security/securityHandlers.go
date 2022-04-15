package security

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// HashData - Uses Sha256 as an encryption algorithm.
func HashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	strData := fmt.Sprintf("%x", hash)
	return strData
}

func ValidateJWT(token string, secret string) (claims jwt.MapClaims, err error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return
	}
	if tok == nil {
		err = errors.New("invalid token string")
		return
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		err = errors.New("invalid signature")
	}
	return
}
