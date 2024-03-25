package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IEnrollment interface {
	CreateEnrollment(req *models.EnrollmentRequest) (*models.EnrollmentResponse, error)
	GetEnrollmentById(id uint) (*models.EnrollmentResponse, error)
	ListEnrollment() ([]models.EnrollmentResponse, error)
	DeleteEnrollment(id uint) error
	UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.EnrollmentResponse, error)
}

func (s Service) CreateEnrollment(req *models.EnrollmentRequest) (*models.EnrollmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewEnrollment(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateEnrollment(newReq)
	if err != nil {
		return nil, err
	}
	return data.EnrollmentResponse(), nil
}

func (s Service) GetEnrollmentById(id uint) (*models.EnrollmentResponse, error) {
	data, err := s.repo.FindEnrollmentById(id)
	if err != nil {
		return nil, err
	}
	return data.EnrollmentResponse(), nil
}

func (s Service) ListEnrollment() ([]models.EnrollmentResponse, error) {
	datas, err := s.repo.FindAllEnrollment()
	if err != nil {
		return nil, err
	}
	var responses []models.EnrollmentResponse
	for _, data := range *datas {
		responses = append(responses, *data.EnrollmentResponse())
	}
	return responses, nil
}

func (s Service) DeleteEnrollment(id uint) error {
	return s.repo.DeleteEnrollment(id)
}

func (s Service) UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.EnrollmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewEnrollment(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateEnrollment(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.EnrollmentResponse(), nil
}
