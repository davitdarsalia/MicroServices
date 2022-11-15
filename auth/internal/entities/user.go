package entities

import "github.com/dgrijalva/jwt-go"

/* User Types */

type User struct {
	UserName  string `json:"username" binding:"required"`
	UserRole  string `json:"user_role"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Country   string `json:"country" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	City      string `json:"city" binding:"required"`
	CreatedAt string `json:"createdat"`
	IpAddress string `json:"ip_address"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshLogin struct {
	RT string `json:"refresh_token" binding:"required"`
}

/* Token Types */

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
