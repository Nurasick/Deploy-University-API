package handler

import (
	"net/http"
	"strconv"
	"university/model"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	ScheduleService service.ScheduleServiceInterface
}

func NewScheduleHandler(scheduleService service.ScheduleServiceInterface) *ScheduleHandler {
	return &ScheduleHandler{ScheduleService: scheduleService}
}

// @Summary Get group schedule
// @Description Retrieve schedule for a specific group
// @Tags Schedules
// @Produce json
// @Security Bearer
// @Param id path int true "Group ID"
// @Success 200 {array} model.Schedule
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /schedule/group/{id} [get]
func (h *ScheduleHandler) GetScheduleByGroupID(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid userID"})
	}

	data, err := h.ScheduleService.GetSchedulesByGroupID(groupID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Get all schedule
// @Description Retrieve all schedules
// @Tags Schedules
// @Produce json
// @Security Bearer
// @Success 200 {array} model.Schedule
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /schedule/all_class_schedule [get]
func (h *ScheduleHandler) GetAllSchedules(c echo.Context) error {

	data, err := h.ScheduleService.GetAllSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// @Summary Create a new class schedule
// @Description Add schedule for a group with teacher and subject
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body model.ScheduleRequest true "Schedule Info"
// @Success 201 {object} model.Schedule
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security Bearer
// @Router /schedule [post]
func (h *ScheduleHandler) CreateSchedule(c echo.Context) error {
	var req model.ScheduleRequest
	if err := c.Bind(&req); err != nil { //binding the request body to struct
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request: " + err.Error()})
	}
	schedule, err := h.ScheduleService.CreateSchedule(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, schedule)

}
