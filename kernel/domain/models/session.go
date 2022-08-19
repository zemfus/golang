package models

import "time"

type Session struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Code      int       `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	EndAt     time.Time `json:"end_at"`
}
