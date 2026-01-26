package handler

import (
	"net/http"
	"strconv"
	"time"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type AttendanceHandler struct {
	AttendanceService service.AttendanceServiceInterface
}

func NewAttendanceHandler(attendanceService service.AttendanceServiceInterface) *AttendanceHandler {
	return &AttendanceHandler{AttendanceService: attendanceService}
}

// @Summary Mark attendance
// @Description Mark attendance for a student in a subject
// @Tags Attendance
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body model.AttendanceRequest true "Attendance Record"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/subject [post]
func (h *AttendanceHandler) MarkAttendance(c echo.Context) error {
	var req model.AttendanceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body" + err.Error()})
	}

	visitDay, err := time.Parse("2006-01-02", req.VisitDay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "visit_day must be YYYY-MM-DD",
		})
	}
	//roleID := c.Get("RoleID").(int)

	err = h.AttendanceService.MarkAttendance(req.StudentID, req.SubjectID, 1, visitDay, req.Visited)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"data": "attendance saved"})
}

// @Summary Get attendance by subject
// @Description Retrieve all attendance records for a specific subject
// @Tags Attendance
// @Produce json
// @Security Bearer
// @Param id path int true "Subject ID"
// @Success 200 {array} model.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/attendanceBySubjectID/{id} [get]
func (h *AttendanceHandler) GetAttendanceBySubjectID(c echo.Context) error {
	subjectID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid subjectID"})
	}

	data, err := h.AttendanceService.GetAttendanceBySubjectID(subjectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Get attendance by student
// @Description Retrieve all attendance records for a specific student
// @Tags Attendance
// @Produce json
// @Security Bearer
// @Param id path int true "Student ID"
// @Success 200 {array} model.AttendanceResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/attendanceByStudentID/{id} [get]
func (h *AttendanceHandler) GetAttendanceByStudentID(c echo.Context) error {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid studentID"})
	}

	data, err := h.AttendanceService.GetAttendanceByStudentID(studentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
