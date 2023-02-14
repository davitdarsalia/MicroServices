package entities

import "time"

type UUID = string
type Email struct {
	SendAt time.Time `json:"send_at"`
	UserID UUID      `json:"user_id"`
}
