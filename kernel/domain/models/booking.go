package models

import (
	"time"
)

type BookType string

const (
	PlacesType    BookType = "places"
	InventoryType BookType = "inventory"
)

type Booking struct {
	ID          int       `json:"id"`
	BookType    BookType  `json:"book_type"`
	UserID      int       `json:"user_id"`
	InventoryID int       `json:"inventory_id"`
	PlacesID    int       `json:"places_id"`
	Confirm     bool      `json:"confirm"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	Status      bool      `json:"status"`
}
