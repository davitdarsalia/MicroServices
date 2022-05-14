package entities

type RegisteredUserResponse struct {
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type SignedInUserResponse struct {
	UserId      int    `json:"user_id"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
