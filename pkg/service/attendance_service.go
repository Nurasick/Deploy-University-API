package service

import (
	"errors"
	"time"
	"university/model"
	"university/pkg/repository"
)

type AttendanceServiceInterface interface {
	MarkAttendance(studentID, subjectID, roleID int, visitDay time.Time, visited bool) error
	GetAttendanceBySubjectID(subjectID int) ([]model.AttendanceResponse, error)
	GetAttendanceByStudentID(subjectID int) ([]model.AttendanceResponse, error)
}

type AttendanceService struct {
	attendanceRepo repository.AttendanceRepositoryInterface
	studentRepo    repository.StudentRepositoryInterface
	subjectRepo    repository.SubjectRepositoryInterface
}

func NewAttendanceService(
	attendanceRepo repository.AttendanceRepositoryInterface,
	studentRepo repository.StudentRepositoryInterface,
	subjectRepo repository.SubjectRepositoryInterface,
) *AttendanceService {
	return &AttendanceService{
		attendanceRepo: attendanceRepo,
		studentRepo:    studentRepo,
		subjectRepo:    subjectRepo,
	}
}

func (r *AttendanceService) MarkAttendance(studentID, subjectID, roleID int, visitDay time.Time, visited bool) error {
	if roleID != 1 && roleID != 2 {
		return errors.New("forbidden: insufficient permissions")
	}
	_, err := r.studentRepo.GetStudentByID(studentID)
	if err != nil {
		return errors.New("student not found: " + err.Error())
	}
	_, err = r.subjectRepo.GetSubjectByID(subjectID)
	if err != nil {
		return errors.New("Subject not found: " + err.Error())
	}
	attendanceExist, err := r.attendanceRepo.Exists(studentID, subjectID, visitDay)
	if err != nil {

		return errors.New("failed to get attendance: " + err.Error())
	}

	if attendanceExist != nil && attendanceExist.ID != 0 {
		return errors.New("attendance already marked for that day")
	}

	attendance := &model.Attendance{
		StudentID: studentID,
		SubjectID: subjectID,
		VisitDay:  visitDay,
		Visited:   visited,
	}
	return r.attendanceRepo.MarkAttendance(attendance)
}

func (r *AttendanceService) GetAttendanceBySubjectID(subjectID int) ([]model.AttendanceResponse, error) {
	records, err := r.attendanceRepo.GetAttendanceBySubjectID(subjectID)

	if err != nil {
		return nil, errors.New("Failed to get records of attendances by subject: " + err.Error())
	}
	var responses []model.AttendanceResponse
	for _, record := range records {
		responses = append(responses, model.AttendanceResponse{
			ID:          record.ID,
			StudentName: record.StudentName,
			SubjectName: record.SubjectName,
			VisitDay:    record.VisitDay.Format("2006-01-02"),
			Visited:     record.Visited,
		})
	}
	return responses, nil
}

func (r *AttendanceService) GetAttendanceByStudentID(subjectID int) ([]model.AttendanceResponse, error) {
	records, err := r.attendanceRepo.GetAttendanceByStudentID(subjectID)

	if err != nil {
		return nil, errors.New("Failed to get records of attendances by subject: " + err.Error())
	}
	var responses []model.AttendanceResponse
	for _, record := range records {
		responses = append(responses, model.AttendanceResponse{
			ID:          record.ID,
			StudentName: record.StudentName,
			SubjectName: record.SubjectName,
			VisitDay:    record.VisitDay.Format("2006-01-02"),
			Visited:     record.Visited,
		})
	}
	return responses, nil
}
