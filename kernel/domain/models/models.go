package models

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
	ID   int `json:"id"`
	Name int `json:"name"`
}

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Campus      int    `json:"campus_id"`
	Category    int    `json:"category_id"`
}
