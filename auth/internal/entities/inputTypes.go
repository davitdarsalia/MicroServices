package entities

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ResetPasswordInput struct {
	Email       string `json:"email" binding:"required"`
	IDNumber    string `json:"id_number" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
