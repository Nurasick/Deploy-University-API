package handler

import (
	"net/http"
	"strconv"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type StudentHandler struct {
	StudentService service.StudentServiceInterface
}

// NewUserHandler creates a new instance of UserHandler
func NewStudentHandler(studentService service.StudentServiceInterface) *StudentHandler {
	return &StudentHandler{StudentService: studentService}
}

// @Summary Get student attendance
// @Description Get attendance records for the current student
// @Tags Student
// @Produce json
// @Security Bearer
// @Param id path int true "Student ID"
// @Success 200 {array} model.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student/{id}/attendance [get]
func (h *StudentHandler) MyAttendance(c echo.Context) error {
	userID := c.Get("userID").(int)

	stud, err := h.StudentService.GetStudentByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	data, err := h.StudentService.GetStudentAttendance(stud.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Get student by ID
// @Description Retrieve student information by user ID
// @Tags Student
// @Produce json
// @Security Bearer
// @Param id path int true "User ID"
// @Success 200 {object} model.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student/{id} [get]
func (h *StudentHandler) GetStudentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}

	stud, err := h.StudentService.GetStudentByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	data := model.StudentResponse{
		ID:        stud.ID,
		Name:      stud.Name,
		BirthDate: stud.BirthDate.Format("2006-01-02"),
		GroupID:   stud.GroupID,
		GroupName: stud.GroupName,
		Gender:    stud.Gender,
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Create a new student
// @Description Create a new student with user and group mapping
// @Tags Student
// @Accept json
// @Produce json
// @Param body body model.StudentRequest true "Student Info"
// @Success 201 {object} model.Student
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /student [post]
func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var req model.StudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	student, err := h.StudentService.CreateStudent(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, student)

}

// @Summary Update an existing student student
// @Description Update an existing user
// @Tags Student
// @Accept json
// @Produce json
// @Param student body model.StudentRequest true "Student Info"
// @Success 201 {object} model.Student
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /student/{id} [patch]
func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	var req model.StudentRequest
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	stud := &model.Student{
		ID:        id,
		Name:      req.Name,
		GroupID:   req.GroupID,
		BirthDate: req.BirthDate,
		UserId:    req.UserId,
		Year:      req.Year,
		Gender:    req.Gender,
	}
	err = h.StudentService.UpdateStudent(stud)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, stud)
}

// @Summary Get all students
// @Description Retrieve students
// @Tags Student
// @Produce json
// @Security Bearer
// @Success 200 {object} []model.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student/all [get]
func (h *StudentHandler) GetAllStudents(c echo.Context) error {
	data, err := h.StudentService.GetAllStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
