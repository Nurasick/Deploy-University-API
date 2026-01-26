package service

import (
	"errors"
	"university/model"
	"university/pkg/repository"
)

type TeacherServiceInterface interface {
	CreateTeacher(teacher *model.TeacherRequest) (*model.Teacher, error)
	GetScheduleByTeacherID(id int) ([]model.Schedule, error)
	GetTeacherByID(id int) (*model.Teacher, error)
}

type TeacherService struct {
	teacherRepo  repository.TeacherRepository
	userRepo     repository.UserRepository
	scheduleRepo repository.ScheduleRepository
}

func NewTeacherService(
	teacherRepo repository.TeacherRepository,
	userRepo repository.UserRepository,
	scheduleRepo repository.ScheduleRepository,
) *TeacherService {
	return &TeacherService{
		teacherRepo:  teacherRepo,
		userRepo:     userRepo,
		scheduleRepo: scheduleRepo,
	}
}

func (r *TeacherService) CreateTeacher(teacher *model.TeacherRequest) (*model.Teacher, error) {
	userID := teacher.UserId
	user, err := r.userRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("User does not exist: " + err.Error())
	}
	if user.RoleID != 2 {
		return nil, errors.New("User is not a teacher")
	}
	teach := &model.Teacher{
		UserId:     teacher.UserId,
		Name:       teacher.Name,
		Department: teacher.Department,
	}
	teacherID, err := r.teacherRepo.CreateTeacher(teach)
	if err != nil {
		return nil, errors.New("Failed to create a teacher: " + err.Error())
	}
	teach.ID = teacherID
	return teach, nil
}

func (r *TeacherService) GetScheduleByTeacherID(id int) ([]model.Schedule, error) {
	teacher, err := r.teacherRepo.GetTeacherByID(id)
	if err != nil {
		return nil, errors.New("Failed to get teacher: " + err.Error())
	}
	schedules, err := r.scheduleRepo.GetScheduleByTeacherID(teacher.ID)
	if err != nil {
		return nil, errors.New("Failed to get schedule by teacher ID: " + err.Error())
	}
	return schedules, nil
}
func (r *TeacherService) GetTeacherByID(id int) (*model.Teacher, error) {
	teacher, err := r.teacherRepo.GetTeacherByID(id)
	if err != nil {
		return nil, errors.New("Failed to get teacher: " + err.Error())
	}
	return teacher, err
}
