package entities

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name        string `json:"name" binding:"required" validate:"required,min=2,max=255"`
	Surname     string `json:"surname" binding:"required" validate:"required,min=2,max=255"`
	UserName    string `json:"username" binding:"required" validate:"required,min=7,max=40"`
	Email       string `json:"email" binding:"required" validate:"required,min=10,max=255,email"`
	TelNumber   string `json:"tel_number" binding:"required" validate:"required,min=5,max=50,e164"`
	IDNumber    string `json:"id_number" binding:"required" validate:"required,min=11,max=11"`
	Password    string `json:"password" binding:"required" validate:"required,min=7,max=200"`
	Salt        string `json:"salt"`
	DateCreated string `json:"date_created"`
	IPAddress   string `json:"ip_address"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required" validate:"required,min=10,max=255,email"`
	Password string `json:"password" binding:"required" validate:"required" validate:"required,min=7,max=200"`
	IDNumber string `json:"id_number" binding:"required" validate:"required" validate:"required,min=11,max=11"`
}

type RecoverPasswordInput struct {
	Email       string `json:"email" binding:"required"`
	IDNumber    string `json:"id_number" binding:"required"`
	TelNumber   string `json:"tel_number" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type AccessToken struct {
	jwt.StandardClaims
	User
	TokenCreatedAtString string `json:"token_created_at"`
}

type AuthenticatedUserResponse struct {
	UserID                string `json:"user_id"`
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresAt  string `json:"access_token_exp"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_exp"`
}
