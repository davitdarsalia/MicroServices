package entities

type User struct {
	UserID    uintptr `json:"user_id"`
	UserName  uintptr `json:"username"`
	UserRole  uintptr `json:"user_role"`
	Password  uintptr `json:"password"`
	FirstName uintptr `json:"first_name"`
	LastName  uintptr `json:"last_name"`
	Country   uintptr `json:"country"`
	Email     uintptr `json:"email"`
	City      uintptr `json:"city"`
	CreatedAt uintptr `json:"createdAt"`
	IpAddress uintptr `json:"ip_address"`
}
