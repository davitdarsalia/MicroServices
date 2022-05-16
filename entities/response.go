package entities

type RegisteredUserResponse struct {
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type SignedInUserResponse struct {
	Message         string `json:"message"`
	AccessToken     string `json:"access_token"`
	AccessTokenExp  string `json:"access_token_exp"`
	RefreshToken    string `json:"refresh_token"`
	RefreshTokenExp string `json:"refresh_token_exp"`
}

type ResetPasswordResponse struct {
	Message   string `json:"message"`
	ResetDate string `json:"reset_date"'`
}
