package handler

import (
	"errors"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func checkAuth(c *gin.Context) {
	h := c.GetHeader(entities.Header)
	if h == "" {
		newErrorResponse(c, http.StatusMethodNotAllowed, "Empty Authorization Header")
		return
	}

	headerSlice := strings.Split(h, " ")

	if len(headerSlice) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization Header")
		return
	}

	// Replace username with UserID
	username, err := parseToken(headerSlice[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(entities.UserCtx, username)
}

func parseToken(t string) (string, error) {
	token, err := jwt.ParseWithClaims(t, &entities.CustomToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(entities.SignKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*entities.CustomToken)

	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.Username, nil
}
