package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IAttendance interface {
	CreateAttendance(req *models.AttendanceRecordRequest) error
	GetAttendanceById(id uint) (*models.AttendanceResponse, error)
	ListAttendance() ([]models.AttendanceResponse, error)
	DeleteAttendance(id uint) error
	UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.AttendanceResponse, error)
}

func (s Service) CreateAttendance(req *models.AttendanceRecordRequest) error {
	if req == nil {
		return errors.New("attendance request is nil while creating attendance")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return err
	}
	_, err := s.repo.FindClassById(req.ClassId)
	if err != nil {
		return errors.New("class not found")
	}
	for _, attendance := range req.Attendance {
		user, err := s.repo.FindUserById(attendance.StudentId)
		if err != nil {
			return errors.New("student not found")
		}
		if user.UserType != "student" {
			return errors.New("student not found")
		}
	}
	for _, attendance := range req.Attendance {
		attendanceRequest := &models.AttendanceRequest{
			Date:      req.Date,
			ClassId:   req.ClassId,
			StudentId: attendance.StudentId,
			Status:    attendance.Status,
		}
		if attendance.Status == "present" || attendance.Status == "absent" {
			attendanceRequest.IsPresent = true
		}
		newReq, err := models.NewAttendance(attendanceRequest)
		if err != nil {
			return err
		}
		_, err = s.repo.CreateAttendance(newReq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetAttendanceById(id uint) (*models.AttendanceResponse, error) {
	data, err := s.repo.FindAttendanceById(id)
	if err != nil {
		return nil, err
	}
	return data.AttendanceResponse(), nil
}

func (s Service) ListAttendance() ([]models.AttendanceResponse, error) {
	datas, err := s.repo.FindAllAttendance()
	if err != nil {
		return nil, err
	}
	var responses []models.AttendanceResponse
	for _, data := range *datas {
		responses = append(responses, *data.AttendanceResponse())
	}
	return responses, nil
}

func (s Service) DeleteAttendance(id uint) error {
	return s.repo.DeleteAttendance(id)
}

func (s Service) UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.AttendanceResponse, error) {
	if req == nil {
		return nil, errors.New("attendance request is nil while updating attendance")
	}
	// req.Prepare()
	// if err := req.Validate(); err != nil {
	// 	return nil, err
	// }
	_, err := s.repo.FindClassById(req.ClassId)
	if err != nil {
		return nil, errors.New("class not found")
	}
	user, err := s.repo.FindUserById(req.StudentId)
	if err != nil {
		return nil, errors.New("student not found")
	}
	if user.UserType != "student" {
		return nil, errors.New("student not found")
	}
	datum, err := s.repo.FindAttendanceById(id)
	if err != nil {
		return nil, errors.New("attendance not found")
	}
	data, err := s.repo.UpdateAttendance(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.AttendanceResponse(), nil
}
