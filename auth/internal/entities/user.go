package entities

import "github.com/dgrijalva/jwt-go"

type User struct {
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	UserName    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	TelNumber   string `json:"tel_number" binding:"required"`
	IDNumber    string `json:"id_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Salt        string `json:"salt"`
	DateCreated string `json:"date_created"`
	IPAddress   string `json:"ip_address"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IDNumber string `json:"id_number" binding:"required"`
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
