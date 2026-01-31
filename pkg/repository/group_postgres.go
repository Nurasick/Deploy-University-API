package repository

import (
	"context"
	"errors"
	"university/model"

	"github.com/jackc/pgx/v5"
)

type GroupRepositoryInterface interface {
	GetAllGroups() ([]model.Group, error)
	CreateGroup(group *model.Group) error
}

type GroupRepository struct {
	conn *pgx.Conn
}

func NewGroupRepository(conn *pgx.Conn) *GroupRepository {
	return &GroupRepository{conn: conn}
}

func (r *GroupRepository) GetAllGroups() ([]model.Group, error) {
	query := `select id,name from groups;`

	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Failed to retrieve groups: " + err.Error())
	}
	defer rows.Close()
	var groups []model.Group
	for rows.Next() {
		var group model.Group
		err := rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, errors.New("Failed to scan group: " + err.Error())
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (r *GroupRepository) CreateGroup(group *model.Group) error {
	query := `insert into groups (name)
	values($1);`
	_, err := r.conn.Exec(context.Background(), query, group.Name)
	if err != nil {
		return errors.New("Failed to create group: " + err.Error())
	}
	return nil
}
