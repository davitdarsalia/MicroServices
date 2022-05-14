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
	UserId   string `json:"user_id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CustomToken struct {
	jwt.StandardClaims
	Role      string `json:"role"`
	Ip        string `json:"ip_address"`
	SessionId string `json:"session_id"`
}
