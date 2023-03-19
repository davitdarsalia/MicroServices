package entities

import "github.com/dgrijalva/jwt-go"

type AccessToken struct {
	jwt.StandardClaims
	AccessTokenCustomClaims
}

type AccessTokenCustomClaims struct {
	TelNumber      string `json:"tel_number"`
	IDNumber       string `json:"id_number"`
	UserID         string `json:"user_id"`
	CreatedAt      string `json:"created_at"`
	UserRole       string `json:"user_role"`
	ExpirationTime string `json:"expiration_time"`
	IpAddress      string `json:"ip_address"`
	// TODO - Add support for client signature for security resons
}
