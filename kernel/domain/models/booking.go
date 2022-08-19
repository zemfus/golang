package models

import (
	"time"
)

type Booking struct {
	ID         int       `json:"id"`
	BookTypeID int       `json:"book_type_id"`
	UserID     int       `json:"user_id"` //
	StartAt    time.Time `json:"start_at"`
	EndAt      time.Time `json:"end_at"`
	Status     bool      `json:"status"`
}
