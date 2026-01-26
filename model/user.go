package model

import "time"

//User model represents a user in the system
type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	RoleID       int       `json:"role_id"`
	Status       int       `json:"status"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	ActiveStatus   = 1
	InactiveStatus = 2
)
