package model

import "time"

type Student struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	GroupID   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	BirthDate time.Time `json:"birth_date"`
	Year      int       `json:"year"`
	Gender    string    `json:"gender"`
	UserId    int       `json:"user_id"`
}

type StudentResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	Gender    string `json:"gender"`
}

type StudentRequest struct {
	Name      string    `json:"name"`
	GroupID   int       `json:"group_id"`
	BirthDate time.Time `json:"birth_date"`
	Year      int       `json:"year"`
	Gender    string    `json:"gender"`
	UserId    int       `json:"user_id"`
}
