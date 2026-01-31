package model

type Schedule struct {
	ID        int    `json:"id"`
	GroupID   int    `json:"group_id"`
	SubjectID int    `json:"subject_id"`
	Day       int    `json:"day_of_week"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	TeacherID int    `json:"teacher_id"`
}

type ScheduleRequest struct {
	GroupID   int    `json:"group_id"`
	SubjectID int    `json:"subject_id"`
	Day       int    `json:"day_of_week"`
	StartsAt  string `json:"starts_at"`
	EndsAt    string `json:"ends_at"`
	TeacherID int    `json:"teacher_id"`
}

type Subject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type GroupRequest struct {
	Name string `json:"name"`
}
