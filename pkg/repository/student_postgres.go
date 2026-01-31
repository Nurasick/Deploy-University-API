package repository

import (
	"context"
	"errors"
	"fmt"
	"university/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepositoryInterface interface {
	CreateStudent(*model.Student) (int, error)
	GetStudentByID(id int) (*model.Student, error)
	GetStudentByUserID(userID int) (*model.Student, error)
	GetAllStudents() ([]model.Student, error)
	UpdateStudent(*model.Student) error
}

type StudentRepository struct {
	conn *pgxpool.Pool
}

func NewStudentRepository(conn *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{conn: conn}
}

func (r *StudentRepository) CreateStudent(student *model.Student) (int, error) {
	query := `insert into students (firstname,surname,birth_date,group_id,gender_id,user_id,year_of_study) 
	values ($1,$2,$3,$4,$5,$6,$7) returning id;`

	var id int
	err := r.conn.QueryRow(context.Background(), query, student.Firstname, student.Surname, student.BirthDate, student.GroupID, student.GenderID, student.UserId, student.Year).Scan(&id)

	if err != nil {
		fmt.Print(student)
		return 0, errors.New("Failed to create student: " + err.Error())
	}

	student.ID = id
	return id, nil

}

func (r *StudentRepository) GetStudentByID(id int) (*model.Student, error) {
	query := `
	SELECT s.id, s.firstname,s.surname, s.birth_date, s.group_id, gr.name, g.name, s.user_id, s.year_of_study
	FROM students s
	JOIN genders g ON g.id = s.gender_id
	join groups gr ON gr.id = s.group_id
	WHERE s.id =$1;`

	var student model.Student

	err := r.conn.QueryRow(context.Background(), query, id).Scan(&student.ID, &student.Firstname, &student.Surname, &student.BirthDate, &student.GroupID, &student.GroupName, &student.Gender, &student.UserId, &student.Year)

	if err != nil {
		return nil, errors.New("Failed to retrieve student: " + err.Error())
	}
	return &student, nil
}

func (r *StudentRepository) GetStudentByUserID(userID int) (*model.Student, error) {
	query := `
	SELECT s.id, s.firstname,s.surname, s.birth_date, s.group_id, gr.name, g.name, s.user_id, s.year_of_study
	FROM students s
	JOIN genders g ON g.id = s.gender_id
	join groups gr ON gr.id = s.group_id
	WHERE s.user_id =$1;`

	var student model.Student

	err := r.conn.QueryRow(context.Background(), query, userID).Scan(&student.ID, &student.Firstname, &student.Surname, &student.BirthDate, &student.GroupID, &student.GroupName, &student.Gender, &student.UserId, &student.Year)
	if err != nil {
		return nil, errors.New("Failed to retrieve student: " + err.Error())
	}
	return &student, nil
}
func (r *StudentRepository) GetAllStudents() ([]model.Student, error) {
	query := `
	SELECT s.id, s.firstname,s.surname, s.birth_date, s.group_id,gr.name, g.name, s.user_id, s.year_of_study
	FROM students s
	JOIN genders g ON g.id = s.gender_id
	join groups gr ON gr.id = s.group_id;`

	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve students: " + err.Error())
	}
	defer rows.Close()
	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Firstname, &student.Surname, &student.BirthDate, &student.GroupID, &student.GroupName, &student.Gender, &student.UserId, &student.Year)
		if err != nil {
			return nil, errors.New("Failed to scan student: " + err.Error())
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *StudentRepository) UpdateStudent(student *model.Student) error {
	query := `UPDATE students SET firstname=$1,surname=$2, birth_date=$3, group_id=$4, gender_id=$5, year=$6 where id = $6;`

	_, err := r.conn.Exec(context.Background(), query, student.Firstname, student.Surname, student.BirthDate, student.GroupID, student.Gender, student.Year, student.ID)
	if err != nil {
		return errors.New("Failed to update student: " + err.Error())
	}
	return nil
}
