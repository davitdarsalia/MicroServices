package entities

import "time"

type Deposit struct {
	UserID                 int           `json:"user_id" binding:"required"`
	Balance                int           `json:"balance" binding:"required"`
	ExpiryDate             time.Duration `json:"expiry_date" binding:"required"`
	AdditionPerTransaction int           `json:"addition_per_transaction" binding:"required"`
	CurrencyID             string        `json:"currency_id" binding:"required"`
}
