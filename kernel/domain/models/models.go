package models

import "time"

type Role string

const (
	Student   Role = "student"
	Staff     Role = "staff"
	Applicant Role = "applicant"
	Dev       Role = "dev"
	Unknown   Role = "unknown"
)

type User struct {
	ID         int    `json:"ID"`
	Nickname   string `json:"nickname"`
	Role       Role   `json:"role"`
	Email      string `json:"email"`
	CampusID   *int   `json:"campus_id"`
	HandleStep int    `json:"handle_step"`
}

type Campus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Places struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CampusID    int           `json:"campus_id"`
	CategoryID  int           `json:"category_id"`
	Floor       int           `json:"floor"`
	Room        int           `json:"room"`
	Period      time.Duration `json:"period"`
	Permission  Role          `json:"permission"`
}

type Inventory struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CampusID    int           `json:"campus_id"`
	CategoryID  int           `json:"category_id"`
	Period      time.Duration `json:"period"`
	Permission  Role          `json:"permission"`
}
