package models

import "time"

type TokenModel struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"expires_in"`
	TokenType    string        `json:"token_type"`
	UserId       string        `json:"user_id"`
}
