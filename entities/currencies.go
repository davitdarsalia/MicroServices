package entities

type Currencies struct {
	Currency string `json:"currency" binding:"required"`
}
