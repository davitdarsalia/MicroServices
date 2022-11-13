package entities

type RegisteredUserResponse struct {
	UserId    uintptr `json:"user_id"`
	Message   string  `json:"message"`
	CreatedAt string  `json:"created_at"`
}
