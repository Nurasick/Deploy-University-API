package handler

import (
	"net/http"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	AuthService *service.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Register handles user registration requests

// @Summary Register a new user
// @Description Create a new user account with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body model.RegisterRequest true "Register Request"
// @Success 201 {object} model.RegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil { //binding the request body to struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	user, err := h.AuthService.Register(req.Email, req.Password, req.RoleID) //RoleID is 1 for now as i do not know how to make them separate and when to pass //calling the Register method of AuthService
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	registerResponse := model.RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return c.JSON(http.StatusCreated, registerResponse)
}

// @Summary Login user
// @Description Login with email and password to get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "Login Request"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error { // Login handles user login requests
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil { //binding the request body to struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	token, err := h.AuthService.Login(req.Email, req.Password) //calling the Login method of AuthService
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, model.AuthResponse{Token: token})
}
