package repository

import (
	"context"
	"errors"
	"university/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherRepositoryInterface interface {
	CreateTeacher(*model.Teacher) (int, error)
	GetTeacherByID(id int) (model.Teacher, error)
	GetTeacherByUserID(userID int) (model.Teacher, error)
	GetAllTeachers() ([]model.Teacher, error)
}

type TeacherRepository struct {
	conn *pgxpool.Pool
}

func NewTeacherRepository(conn *pgxpool.Pool) *TeacherRepository {
	return &TeacherRepository{conn: conn}
}

func (r *TeacherRepository) CreateTeacher(teacher *model.Teacher) (int, error) {
	query := `insert into teachers (full_name,department,user_id)
	values ($1,$2,$3) returning id;`

	var id int
	err := r.conn.QueryRow(context.Background(), query, teacher.Name, teacher.Department, teacher.UserId).Scan(&id)
	if err != nil {
		return 0, errors.New("Failed to create a teacher: " + err.Error())
	}
	return id, nil
}

func (r *TeacherRepository) GetTeacherByID(id int) (*model.Teacher, error) {
	query := `select id, full_name, department, user_id from teachers where id=$1;`

	var teacher model.Teacher
	err := r.conn.QueryRow(context.Background(), query, id).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Department,
		&teacher.UserId,
	)
	if err != nil {
		return nil, errors.New("Failed to get a teacher by ID: " + err.Error())
	}
	return &teacher, nil
}

func (r *TeacherRepository) GetTeacherByUserID(userID int) (*model.Teacher, error) {
	query := `select id, full_name, department, user_id from teachers where user_id=$1;`
	var teacher model.Teacher
	err := r.conn.QueryRow(context.Background(), query, userID).Scan(
		&teacher.ID,
		&teacher.Name,
		&teacher.Department,
		&teacher.UserId,
	)
	if err != nil {
		return nil, errors.New("Failed to get a teacher by userID: " + err.Error())
	}
	return &teacher, nil

}

func (r *TeacherRepository) GetAllTeachers() ([]model.Teacher, error) {
	query := `select id, fuLL_name, department, user_id from teachers`
	var teachers []model.Teacher
	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve teachers: " + err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var teacher model.Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Department, &teacher.UserId)
		if err != nil {
			return nil, errors.New("Failed to scan teacher: " + err.Error())
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}
