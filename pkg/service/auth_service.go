package service

import (
	"errors"
	"fmt"
	"net/mail"
	"university/helpers/jwt"
	"university/model"
	"university/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(email, password, username string, roleID int) (*model.User, error)
	Login(email, password string) (string, error)
}

// AuthService handles authentication-related operations
type AuthService struct {
	AuthRepo repository.UserRepositoryInterface
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(authRepo *repository.UserRepository) *AuthService {
	return &AuthService{AuthRepo: authRepo}
}

// Register creates a new user with the given email and password. Returns the created user or an error.
func (s *AuthService) Register(email, password string, roleID int) (*model.User, error) {
	_, err := mail.ParseAddress(email) //validating email
	if err != nil {
		return nil, err
	}
	user, err := s.AuthRepo.GetUserByEmail(email) //checking if user already exists
	if err == nil && user != nil {
		return nil, errors.New("Already existing user")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //encrypting the password
	if err != nil {
		return nil, err
	}
	var newUser = &model.User{
		Email:        email,
		PasswordHash: string(passwordHash),
		RoleID:       roleID,
		Status:       model.ActiveStatus,
	}
	userID, err := s.AuthRepo.CreateUser(newUser)
	if err != nil {
		return nil, errors.New("Failed to create a user" + err.Error())
	}
	newUser.ID = userID
	return newUser, nil
}

// Login authenticates a user with the given email and password. Should return a jwt token if success
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.AuthRepo.GetUserByEmail(email) //fetching user by email
	if err != nil && user == nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) //comparing given password with stored password hash
	if err != nil {
		fmt.Println("Password mismatch:", err, user.Email, user.PasswordHash, password)
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(user.ID, user.RoleID, user.Email) // if password matches generating token
	if err != nil {
		return "", err
	}

	return token, nil

}
