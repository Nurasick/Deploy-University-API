package service

import (
	"university/model"
	"university/pkg/repository"
)

type UserServiceInterface interface {
	GetCurrentUser(id int) (*model.MeResponse, error)
	UpdateCurrentUser(user *model.User) error
	DeactivateCurrentUser(id int) error
}

// UserService handles user-related operations
type UserService struct {
	UserRepo repository.UserRepositoryInterface
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// GetCurrentUser retrieves the current user's information by their ID.
func (r *UserService) GetCurrentUser(id int) (*model.MeResponse, error) {
	user, err := r.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &model.MeResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.RoleID,
	}, nil
}

func (r *UserService) UpdateCurrentUser(user *model.User) error {
	return r.UserRepo.UpdateUser(user)
}

func (r *UserService) DeactivateCurrentUser(id int) error {
	return r.UserRepo.DeactivateUser(id)
}
