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
