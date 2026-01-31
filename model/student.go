package model

import "time"

type Student struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	Surname   string    `json:"surname"`
	GroupID   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	BirthDate time.Time `json:"birth_date"`
	Year      int       `json:"year"`
	Gender    string    `json:"gender"`
	GenderID  int       `json:"gender_id"`
	UserId    int       `json:"user_id"`
}

type StudentResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birth_date"`
	Year      int    `json:"year"`
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	Gender    string `json:"gender"`
}

type StudentRequest struct {
	Firstname string    `json:"firstname"`
	Surname   string    `json:"surname"`
	GroupID   int       `json:"group_id"`
	BirthDate time.Time `json:"birth_date"`
	Year      int       `json:"year"`
	GenderID  int       `json:"gender_id"`
	UserId    int       `json:"user_id"`
}
