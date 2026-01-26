package model

type Teacher struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	UserId     int    `json:"user_id"`
}
type TeacherRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	UserId     int    `json:"user_id"`
}
