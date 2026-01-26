package handler

import (
	"net/http"
	"strconv"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type TeacherHandler struct {
	TeacherService    service.TeacherServiceInterface
	AttendanceService service.AttendanceServiceInterface
}

func NewTeacherHandler(
	teacherService service.TeacherServiceInterface,
	attendanceService service.AttendanceServiceInterface,
) *TeacherHandler {
	return &TeacherHandler{
		TeacherService:    teacherService,
		AttendanceService: attendanceService,
	}
}

// @Summary Create a new teacher
// @Description Add a teacher and link to a user
// @Tags Teacher
// @Accept json
// @Produce json
// @Param body body model.TeacherRequest true "Teacher Info"
// @Success 201 {object} model.Teacher
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /teacher [post]
func (h *TeacherHandler) CreateTeacher(c echo.Context) error {
	var req model.TeacherRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request: " + err.Error()})
	}
	teacherID, err := h.TeacherService.CreateTeacher(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, teacherID)
}

// @Summary Get teacher by id
// @Description Get teacher by its id
// @Tags Teacher
// @Produce json
// @Security Bearer
// @Param id path int true "Teacher ID"
// @Success 200 {object} model.Teacher
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teacher/{id} [get]
func (h *TeacherHandler) GetTeacherByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}

	teacher, err := h.TeacherService.GetTeacherByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, teacher)
}
