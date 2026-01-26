package handler

import (
	//"net/http"
	//"university/model"
	"university/pkg/service"
	//"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	UserService     service.UserServiceInterface
	StudentService  service.StudentServiceInterface
	TeacherService  service.TeacherServiceInterface
	ScheduleService service.ScheduleServiceInterface
}

// NewUserHandler creates a new instance of UserHandler
func NewAdminHandler(
	userService service.UserServiceInterface,
	studentService service.StudentServiceInterface,
	teacherService service.TeacherServiceInterface,
	scheduleService service.ScheduleServiceInterface,
) *AdminHandler {
	return &AdminHandler{
		UserService:     userService,
		StudentService:  studentService,
		TeacherService:  teacherService,
		ScheduleService: scheduleService,
	}
}

/*
func (h *AdminHandler) CreateStudent(c echo.Context) error {
	var req model.Student
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	student, err := h.StudentService.CreateStudent(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": student})
}
*/

/*
func (h *AdminHandler) CreateTeacher(c echo.Context) error {
	var req model.Teacher
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	teacher, err := h.TeacherService.CreateTeacher(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": teacher})
}
*/
/*
func (h *AdminHandler) CreateSubject(c echo.Context) error {
	var req model.Subject
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	schedule, err := h.ScheduleService.CreateSchedule(&model.Schedule{
		SubjectID: req.ID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": schedule})
}
*/
