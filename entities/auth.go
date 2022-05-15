package entities

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserID         int    `json:"user_id"`
	PersonalNumber string `json:"personal_number" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	UserName       string `json:"user_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	IpAddress      string `json:"ip_address" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Salt           []byte `json:"salt"`
}

type UserInput struct {
	UserId   string `json:"user_id" db:"user_id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CustomToken struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Ip       string `json:"ip_address"`
}

const (
	Header  = "Authorization"
	SignKey = "466785cf408836b1f39aea588291b9aef5838439c65833f4c1bf6d09022802ba"
	UserCtx = "userName"
)
