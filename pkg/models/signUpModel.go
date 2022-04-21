package models

import "time"

type UserSignUp struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int8   `json:"age"`
	Password  string `json:"password"`
}

type CheckUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenModel struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	TokenType    string        `json:"token_type"`
	UserId       string        `json:"user_id"`
}
