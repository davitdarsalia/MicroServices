package entities

import "github.com/dgrijalva/jwt-go"

type AccessToken struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Ip     string `json:"ip_address"`
}

type RefreshToken struct {
	jwt.StandardClaims
	Ip string `json:"ip_address"`
}
