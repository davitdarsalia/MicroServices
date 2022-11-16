package entities

type RegisteredUser struct {
	UserId    string `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type Authenticated struct {
	AT    string `json:"access_token"`
	ATExp string `json:"access_token_expiry"`
	RT    string `json:"refresh_token"`
	RTExp string `json:"refresh_token_expiry"`
}

type RegisteredResponse struct {
	RegisteredUser
	Authenticated
}

type LoggedInUserResponse struct {
	Message string `json:"message"`
	Authenticated
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}
