package service

import (
	"errors"
	"fmt"
	"strings"
	"university/model"
	"university/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type StudentServiceInterface interface {
	CreateStudent(student *model.StudentRequest) (*model.Student, error)
	GetStudentByID(id int) (*model.Student, error)
	GetStudentByUserID(id int) (*model.Student, error)
	GetStudentAttendance(studentID int) ([]model.AttendanceResponse, error)
	GetAllStudents() ([]model.Student, error)
	UpdateStudent(student *model.Student) error
}

type StudentService struct {
	studentRepo    repository.StudentRepositoryInterface
	userRepo       repository.UserRepositoryInterface
	attendanceRepo repository.AttendanceRepositoryInterface
	authRepo       repository.UserRepositoryInterface
}

func NewStudentService(
	studentRepo *repository.StudentRepository,
	userRepo *repository.UserRepository,
	attendanceRepo *repository.AttendanceRepository,
	authRepo *repository.UserRepository,
) *StudentService {
	return &StudentService{
		studentRepo:    studentRepo,
		userRepo:       userRepo,
		attendanceRepo: attendanceRepo,
		authRepo:       authRepo,
	}
}

func (r *StudentService) CreateStudent(student *model.StudentRequest) (*model.Student, error) {
	email := fmt.Sprintf("%s.%s@university.com",
		strings.ToLower(student.Firstname),
		strings.ToLower(student.Surname),
	)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Email:        email,
		PasswordHash: string(passwordHash),
		RoleID:       3,
		Status:       model.ActiveStatus,
	}
	id, err := r.authRepo.CreateUser(user)
	if err != nil {
		return nil, errors.New("Failed to create a user" + err.Error())
	}

	stud := &model.Student{
		Firstname: student.Firstname,
		Surname:   student.Surname,
		GroupID:   student.GroupID,
		GenderID:  student.GenderID,
		BirthDate: student.BirthDate,
		Year:      student.Year,
		UserId:    id,
	}
	studentID, err := r.studentRepo.CreateStudent(stud)
	if err != nil {
		return nil, errors.New("Failed to create student: " + err.Error())
	}
	stud.ID = studentID
	return stud, nil
}

func (r *StudentService) GetStudentByID(id int) (*model.Student, error) {
	student, err := r.studentRepo.GetStudentByID(id)
	if err != nil {
		return nil, errors.New("Failed to get student: " + err.Error())
	}
	return student, nil
}

func (r *StudentService) GetStudentByUserID(id int) (*model.Student, error) {
	student, err := r.studentRepo.GetStudentByUserID(id)
	if err != nil {
		return nil, errors.New("Failed to get student: " + err.Error())
	}
	return student, nil
}

func (r *StudentService) GetStudentAttendance(studentID int) ([]model.AttendanceResponse, error) {
	attendances, err := r.attendanceRepo.GetAttendanceByStudentID(studentID)
	if err != nil {
		return nil, errors.New("Failed to get student attendance: " + err.Error())
	}
	var attendanceResponses []model.AttendanceResponse
	for _, attendance := range attendances {
		attendanceResponses = append(attendanceResponses, model.AttendanceResponse{
			ID:          attendance.ID,
			StudentName: attendance.StudentName,
			SubjectName: attendance.SubjectName,
			VisitDay:    attendance.VisitDay.Format("2006-01-02"),
			Visited:     attendance.Visited,
		})
	}
	return attendanceResponses, nil
}

func (r *StudentService) UpdateStudent(student *model.Student) error {
	err := r.studentRepo.UpdateStudent(student)
	if err != nil {
		return errors.New("Failed to update student: " + err.Error())
	}
	return nil
}
func (r *StudentService) GetAllStudents() ([]model.Student, error) {
	students, err := r.studentRepo.GetAllStudents()
	if err != nil {
		return nil, errors.New("Failed to get all students: " + err.Error())
	}
	return students, nil
}
