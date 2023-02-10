package responses

import "menuAPI/internal/entities"

type CreateUserResponse struct {
	UserID string `json:"user_id"`
	CreateUserGenericMessage
	entities.AuthenticatedUserResponse
}

type CreateUserGenericMessage struct {
	StatusCode int16  `json:"status_code"`
	Message    string `json:"message"`
}

// Login User

type LoginUserResponse struct {
	UserID string `json:"user_id"`
	LoginUserGenericMessage
	entities.AuthenticatedUserResponse
}

type LoginUserGenericMessage struct {
	// Access Tokens here
	StatusCode int16  `json:"status_code"`
	Message    string `json:"message"`
}

type RecoveredPasswordResponse struct {
	StatusCode int16  `json:"status_code"`
	Message    string `json:"message"`
}
