package repository

import (
	"context"
	"errors"
	"time"
	"university/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepositoryInterface interface {
	MarkAttendance(attendance *model.Attendance) error
	GetAttendanceByStudentID(studentID int) ([]model.Attendance, error)
	GetAttendanceBySubjectID(subjectID int) ([]model.Attendance, error)
	GetAttendanceByGroupIDandDate(groupID int, date time.Time) ([]model.Attendance, error)
	Exists(studentID, subjectID int, visitDay time.Time) (*model.Attendance, error)
}

type AttendanceRepository struct {
	conn *pgxpool.Pool
}

func NewAttendanceRepository(conn *pgxpool.Pool) *AttendanceRepository {
	return &AttendanceRepository{conn: conn}
}

func (r *AttendanceRepository) MarkAttendance(attendance *model.Attendance) error {
	query := `
	insert into attendance(student_id, subject_id, visit_day, visited)
	values($1, $2, $3, $4);
	`

	_, err := r.conn.Exec(context.Background(), query, attendance.StudentID, attendance.SubjectID, attendance.VisitDay, attendance.Visited)
	if err != nil {
		return errors.New("Failed to mark attendance: " + err.Error())
	}
	return nil
}

func (r *AttendanceRepository) GetAttendanceByStudentID(studentID int) ([]model.Attendance, error) {
	query := `
	select a.id, a.student_id, s.name, a.subject_id, sub.name, a.visit_day, a.visited from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.student_id = $1;
	`
	rows, err := r.conn.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, errors.New("Failed to get attendance by student ID: " + err.Error())
	}
	defer rows.Close()

	var attendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.StudentName, &attendance.SubjectID, &attendance.SubjectName, &attendance.VisitDay, &attendance.Visited)
		if err != nil {
			return nil, errors.New("Failed to scan attendance row: " + err.Error())
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}
func (r *AttendanceRepository) GetAttendanceBySubjectID(studentID int) ([]model.Attendance, error) {
	query := `
	select a.id, a.student_id, s.name, a.subject_id, sub.name, a.visit_day, a.visited from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.subject_id = $1;
	`
	rows, err := r.conn.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, errors.New("Failed to get attendance by student ID: " + err.Error())
	}
	defer rows.Close()

	var attendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.StudentName, &attendance.SubjectID, &attendance.SubjectName, &attendance.VisitDay, &attendance.Visited)
		if err != nil {
			return nil, errors.New("Failed to scan attendance row: " + err.Error())
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}
func (r *AttendanceRepository) GetAttendanceByGroupIDandDate(groupID int, date time.Time) ([]model.Attendance, error) {
	query := `
	select a.id, a.student_id, s.name, a.subject_id, sub.name, a.visit_day, a.visited from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.group_id = $1 and a.visit_day = $2;
	`
	rows, err := r.conn.Query(context.Background(), query, groupID, date)
	if err != nil {
		return nil, errors.New("Failed to get attendance by student ID: " + err.Error())
	}
	defer rows.Close()

	var attendances []model.Attendance
	for rows.Next() {
		var attendance model.Attendance
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.StudentName, &attendance.SubjectID, &attendance.SubjectName, &attendance.VisitDay, &attendance.Visited)
		if err != nil {
			return nil, errors.New("Failed to scan attendance row: " + err.Error())
		}
		attendances = append(attendances, attendance)
	}

	return attendances, nil
}

func (r *AttendanceRepository) Exists(studentID, subjectID int, visitDay time.Time) (*model.Attendance, error) {
	query := `
	select a.id, a.student_id, s.name, a.subject_id, sub.name, a.visit_day, a.visited from attendance a
	join students s on a.student_id = s.id
	join subjects sub on a.subject_id = sub.id
	where a.student_id=$1 and a.subject_id = $2 and a.visit_day = $3;
	`
	var attendance model.Attendance
	err := r.conn.QueryRow(context.Background(), query, studentID, subjectID, visitDay).Scan(&attendance.ID, &attendance.StudentID, &attendance.StudentName, &attendance.SubjectID, &attendance.SubjectName, &attendance.VisitDay, &attendance.Visited)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.New("Failed to get attendance by several params: " + err.Error())
	}
	return &attendance, nil
}
