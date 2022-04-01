package models

type UserSignUp struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Age        int8   `json:"age"`
	SecretWord string `json:"secretword"`
}
