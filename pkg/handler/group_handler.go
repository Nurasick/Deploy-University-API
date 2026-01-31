package handler

import (
	"net/http"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type GroupHandler struct {
	GroupService service.GroupServiceInterface
}

// NewUserHandler creates a new instance of UserHandler
func NewGroupHandler(groupService service.GroupServiceInterface) *GroupHandler {
	return &GroupHandler{GroupService: groupService}
}

// @Summary Get all groups
// @Description Retrieve groups
// @Tags Group
// @Produce json
// @Security Bearer
// @Success 200 {object} []model.Group
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /groups [get]
func (h *GroupHandler) GetAllStudents(c echo.Context) error {
	data, err := h.GroupService.GetAllGroups()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}

// @Summary Create a new group
// @Description Create a new group
// @Tags Group
// @Accept json
// @Produce json
// @Param body body model.GroupRequest true "Group Info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /groups [post]
func (h *GroupHandler) CreateGroup(c echo.Context) error {
	var req model.GroupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	gr := &model.Group{
		Name: req.Name,
	}
	err := h.GroupService.CreateGroup(gr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	return c.JSON(http.StatusCreated, "Group Created")

}
