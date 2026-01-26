package handler

import (
	"net/http"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	UserService service.UserServiceInterface
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Me retrieves the current user's information

// @Summary Get current user information
// @Description Get the authenticated user's profile information
// @Tags Users
// @Produce json
// @Security Bearer
// @Success 200 {object} model.MeResponse
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/me [get]
func (h *UserHandler) Me(c echo.Context) error {
	idVal := c.Get("userID") //getting user ID from context
	id, ok := idVal.(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized(userId not valid)"})
	}
	user, err := h.UserService.GetCurrentUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user" + err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}
