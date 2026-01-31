package repository

import (
	"context"
	"errors"
	"university/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ScheduleRepositoryInterface interface {
	CreateSchedule(*model.Schedule) (int, error)
	GetAllSchedule() ([]model.Schedule, error)
	GetScheduleByGroupID(groupID int) ([]model.Schedule, error)
	GetScheduleByTeacherID(teacherID int) ([]model.Schedule, error)
	DeleteSchedule(id int) error
}

type ScheduleRepository struct {
	conn *pgxpool.Pool
}

// NewUserRepository creates a new instance of UserRepository
func NewScheduleRepository(conn *pgxpool.Pool) *ScheduleRepository {
	return &ScheduleRepository{conn: conn}
}

func (r *ScheduleRepository) CreateSchedule(schedule *model.Schedule) (int, error) {
	query := `
	insert into class_schedule(group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id)
	values($1, $2, $3, $4, $5, $6)returning id;
	`
	var id int

	err := r.conn.QueryRow(context.Background(), query, schedule.GroupID, schedule.SubjectID, schedule.Day, schedule.StartsAt, schedule.EndsAt, schedule.TeacherID).Scan(&id)

	if err != nil {
		return 0, errors.New("Failed to create schedule: " + err.Error())
	}
	return id, nil
}

func (r *ScheduleRepository) GetScheduleByGroupID(groupID int) ([]model.Schedule, error) {
	query := `
	select id, group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id from class_schedule where group_id = $1;
	`
	rows, err := r.conn.Query(context.Background(), query, groupID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedule by group ID: " + err.Error())
	}
	defer rows.Close()
	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.Day, &schedule.StartsAt, &schedule.EndsAt, &schedule.TeacherID)
		if err != nil {
			return nil, errors.New("Failed to scan schedule by group ID: " + err.Error())
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (r *ScheduleRepository) GetScheduleByTeacherID(teacherID int) ([]model.Schedule, error) {
	query := `
	select id, group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id from class_schedule where teacher_id = $1;
	`
	rows, err := r.conn.Query(context.Background(), query, teacherID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedule by teacher ID: " + err.Error())
	}
	defer rows.Close()
	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.Day, &schedule.StartsAt, &schedule.EndsAt, &schedule.TeacherID)
		if err != nil {
			return nil, errors.New("Failed to scan schedule by teacher ID: " + err.Error())
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (r *ScheduleRepository) DeleteSchedule(id int) error {
	query := `
	delete from class_schedule where id = $1;
	`
	_, err := r.conn.Exec(context.Background(), query, id)
	if err != nil {
		return errors.New("Failed to delete schedule: " + err.Error())
	}
	return nil
}

func (r *ScheduleRepository) GetAllSchedule() ([]model.Schedule, error) {
	query := `
	select id, group_id, subject_id, day_of_week, starts_at, ends_at, teacher_id from class_schedule;
	`
	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedules: " + err.Error())
	}
	defer rows.Close()
	var schedules []model.Schedule
	for rows.Next() {
		var schedule model.Schedule
		err := rows.Scan(&schedule.ID, &schedule.GroupID, &schedule.SubjectID, &schedule.Day, &schedule.StartsAt, &schedule.EndsAt, &schedule.TeacherID)
		if err != nil {
			return nil, errors.New("Failed to scan schedule: " + err.Error())
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}
