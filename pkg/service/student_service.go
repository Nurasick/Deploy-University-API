package service

import (
	"errors"
	"university/model"
	"university/pkg/repository"
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
}

func NewStudentService(
	studentRepo *repository.StudentRepository,
	userRepo *repository.UserRepository,
	attendanceRepo *repository.AttendanceRepository,
) *StudentService {
	return &StudentService{
		studentRepo:    studentRepo,
		userRepo:       userRepo,
		attendanceRepo: attendanceRepo,
	}
}

func (r *StudentService) CreateStudent(student *model.StudentRequest) (*model.Student, error) {
	userID := student.UserId
	user, err := r.userRepo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("User does not exist: " + err.Error())
	}
	if user.RoleID != 3 {
		return nil, errors.New("User is not a student: ")
	}
	stud := &model.Student{
		Firstname: student.Firstname,
		Surname:   student.Surname,
		GroupID:   student.GroupID,
		Gender:    student.Gender,
		BirthDate: student.BirthDate,
		Year:      student.Year,
		UserId:    student.UserId,
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
