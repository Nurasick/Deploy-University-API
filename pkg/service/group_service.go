package service

import (
	"errors"
	"university/model"
	"university/pkg/repository"
)

type GroupServiceInterface interface {
	CreateGroup(group *model.Group) error
	GetAllGroups() ([]model.Group, error)
}

type GroupService struct {
	GroupRepo repository.GroupRepositoryInterface
}

func NewGroupService(
	groupRepo *repository.GroupRepository,
) *GroupService {
	return &GroupService{
		GroupRepo: groupRepo,
	}
}
func (r *GroupService) CreateGroup(group *model.Group) error {
	err := r.GroupRepo.CreateGroup(group)
	if err != nil {
		return errors.New("Couldn't create group: " + err.Error())
	}
	return nil
}

func (r *GroupService) GetAllGroups() ([]model.Group, error) {
	groups, err := r.GroupRepo.GetAllGroups()
	if err != nil {
		return nil, errors.New("Couldn't get groups: " + err.Error())
	}
	return groups, nil

}
