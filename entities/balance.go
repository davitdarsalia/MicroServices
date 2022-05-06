package entities

type Balance struct {
	UserID         int `json:"user_id" binding:"required"`
	OverallBalance int `json:"overall_balance" binding:"required"`
	Debts          int `json:"debts" binding:"required"`
}
