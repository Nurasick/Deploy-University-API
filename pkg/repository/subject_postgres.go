package repository

import (
	"context"
	"errors"
	"university/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SubjectRepositoryInterface interface {
	CreateSubject(*model.Subject) error
	GetSubjectByID(int) (*model.Subject, error)
	GetAllSubjects() ([]model.Subject, error)
}
type SubjectRepository struct {
	conn *pgxpool.Pool
}

func NewSubjectRepository(conn *pgxpool.Pool) *SubjectRepository {
	return &SubjectRepository{conn: conn}
}

func (r *SubjectRepository) CreateSubject(subject *model.Subject) error {
	var id int
	query := `
	insert into subjects(name)
	values($1) returning id
	`
	err := r.conn.QueryRow(context.Background(), query, subject.Name).Scan(&id)
	if err != nil {
		return errors.New("Failed to create a subject: " + err.Error())
	}

	return nil
}

func (r *SubjectRepository) GetSubjectByID(id int) (*model.Subject, error) {
	query := `select id, name from subjects where id=$1;`

	var subject model.Subject
	err := r.conn.QueryRow(context.Background(), query, id).Scan(
		&subject.ID, &subject.Name,
	)
	if err != nil {
		return &model.Subject{}, errors.New("Failed to get a subject by ID: " + err.Error())
	}
	return &subject, nil
}

func (r *SubjectRepository) GetAllSubjects() ([]model.Subject, error) {
	query := `select id, name from subjects;`

	var subjects []model.Subject
	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve subjects: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var subject model.Subject
		err := rows.Scan(&subject.ID, &subject.Name)
		if err != nil {
			return nil, errors.New("Failed to scan subjects " + err.Error())
		}
		subjects = append(subjects, subject)
	}

	return subjects, nil
}
