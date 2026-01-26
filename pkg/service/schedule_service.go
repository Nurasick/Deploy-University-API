package service

import (
	"errors"
	"university/model"
	"university/pkg/repository"
)

type ScheduleServiceInterface interface {
	CreateSchedule(schedule *model.ScheduleRequest) (*model.Schedule, error)
	GetSchedulesByGroupID(groupID int) ([]model.Schedule, error)
	GetAllSchedule() ([]model.Schedule, error)
}

type ScheduleService struct {
	ScheduleRepo repository.ScheduleRepositoryInterface
}

func NewScheduleService(
	scheduleRepo *repository.ScheduleRepository,
) *ScheduleService {
	return &ScheduleService{
		ScheduleRepo: scheduleRepo,
	}
}

func (r *ScheduleService) CreateSchedule(scheduleRequest *model.ScheduleRequest) (*model.Schedule, error) {
	if scheduleRequest.Day < 0 || scheduleRequest.Day > 7 {
		return nil, errors.New("Invalid day of the week")
	}

	schedule := &model.Schedule{
		SubjectID: scheduleRequest.SubjectID,
		GroupID:   scheduleRequest.GroupID,
		Day:       scheduleRequest.Day,
		StartsAt:  scheduleRequest.StartsAt,
		EndsAt:    scheduleRequest.EndsAt,
		TeacherID: scheduleRequest.TeacherID,
	}
	scheduleID, err := r.ScheduleRepo.CreateSchedule(schedule)
	if err != nil {
		return nil, errors.New("Failed to create schedule: " + err.Error())
	}
	schedule.ID = scheduleID
	return schedule, nil
}

func (r *ScheduleService) GetSchedulesByGroupID(groupID int) ([]model.Schedule, error) {
	schedules, err := r.ScheduleRepo.GetScheduleByGroupID(groupID)
	if err != nil {
		return nil, errors.New("Failed to retrieve schedules by groupID: " + err.Error())
	}
	return schedules, nil

}
func (r *ScheduleService) GetAllSchedule() ([]model.Schedule, error) {
	schedules, err := r.ScheduleRepo.GetAllSchedule()

	if err != nil {
		return nil, errors.New("Failed to retrieve schedules: " + err.Error())
	}
	return schedules, nil
}
