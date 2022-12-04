package entities

type User struct {
	UserName  string `json:"username" binding:"required"`
	UserRole  string `json:"user_role"  binding:"required"`
	Password  string `json:"password"  binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Country   string `json:"country" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	City      string `json:"city" binding:"required"`
	CreatedAt string `json:"createdat"`
	IpAddress string `json:"ip_address"`
}

type RefreshLogin struct {
	RT string `json:"refresh_token" binding:"required"`
}
