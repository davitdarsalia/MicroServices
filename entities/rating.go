package entities

type Rating struct {
	UserID int `json:"user_id" binding:"required"`
	Rating int `json:"rating" binding:"required"`
}
