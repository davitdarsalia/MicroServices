package models

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
