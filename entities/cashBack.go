package entities

type Cashback struct {
	UserID          int `json:"user_id" binding:"required"`
	OverallCashback int `json:"overall_cashback" binding:"required"`
}
