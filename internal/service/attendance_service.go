package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IAttendance interface {
	CreateAttendance(req *models.AttendanceRequest) (*models.AttendanceResponse, error)
	GetAttendanceById(id uint) (*models.AttendanceResponse, error)
	ListAttendance() ([]models.AttendanceResponse, error)
	DeleteAttendance(id uint) error
	UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.AttendanceResponse, error)
}

func (s Service) CreateAttendance(req *models.AttendanceRequest) (*models.AttendanceResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewAttendance(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateAttendance(newReq)
	if err != nil {
		return nil, err
	}
	return data.AttendanceResponse(), nil
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
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewAttendance(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateAttendance(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.AttendanceResponse(), nil
}
