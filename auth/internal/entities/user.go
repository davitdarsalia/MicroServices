package entities

type User struct {
	Name        string `json:"name" binding:"required" validate:"required,min=2,max=255"`
	Surname     string `json:"surname" binding:"required" validate:"required,min=2,max=255"`
	UserName    string `json:"username" binding:"required" validate:"required,min=7,max=40"`
	Email       string `json:"email" binding:"required" validate:"required,min=10,max=255,email"`
	TelNumber   string `json:"tel_number" binding:"required" validate:"required,min=5,max=50,e164"`
	IDNumber    string `json:"id_number" binding:"required" validate:"required,min=11,max=11"`
	Password    string `json:"password" binding:"required" validate:"required,min=7,max=200"`
	Salt        string `json:"salt"`
	DateCreated string `json:"date_created"`
	IPAddress   string `json:"ip_address"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required" validate:"required,min=10,max=255,email"`
	Password string `json:"password" binding:"required"  validate:"required,min=7,max=200"`
	IDNumber string `json:"id_number" binding:"required" validate:"required,min=11,max=11"`
}

type RecoverPasswordInput struct {
	Email       string `json:"email" binding:"required" validate:"required,min=10,max=255,email"`
	IDNumber    string `json:"id_number" binding:"required" validate:"required,min=11,max=11"`
	TelNumber   string `json:"tel_number" binding:"required" validate:"required,min=5,max=50,e164"`
	NewPassword string `json:"new_password" binding:"required" validate:"required,min=7,max=200"`
}

type AuthenticatedUserResponse struct {
	UserID                string `json:"user_id"`
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresAt  string `json:"access_token_exp"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_exp"`
}

type UserMetaInfo struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
	UserID   string `json:"user_id"`
}
