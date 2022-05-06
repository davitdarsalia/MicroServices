package entities

type Bonus struct {
	UserID       int `json:"user_id" binding:"required"`
	OverallBonus int `json:"overall_bonus" binding:"required"`
}
