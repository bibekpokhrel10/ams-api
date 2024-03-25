package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type ISemester interface {
	CreateSemester(req *models.SemesterRequest) (*models.SemesterResponse, error)
	GetSemesterById(id uint) (*models.SemesterResponse, error)
	ListSemester() ([]models.SemesterResponse, error)
	DeleteSemester(id uint) error
	UpdateSemester(id uint, req *models.SemesterRequest) (*models.SemesterResponse, error)
}

func (s Service) CreateSemester(req *models.SemesterRequest) (*models.SemesterResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewSemester(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateSemester(newReq)
	if err != nil {
		return nil, err
	}
	return data.SemesterResponse(), nil
}

func (s Service) GetSemesterById(id uint) (*models.SemesterResponse, error) {
	data, err := s.repo.FindSemesterById(id)
	if err != nil {
		return nil, err
	}
	return data.SemesterResponse(), nil
}

func (s Service) ListSemester() ([]models.SemesterResponse, error) {
	datas, err := s.repo.FindAllSemester()
	if err != nil {
		return nil, err
	}
	var responses []models.SemesterResponse
	for _, data := range *datas {
		responses = append(responses, *data.SemesterResponse())
	}
	return responses, nil
}

func (s Service) DeleteSemester(id uint) error {
	return s.repo.DeleteSemester(id)
}

func (s Service) UpdateSemester(id uint, req *models.SemesterRequest) (*models.SemesterResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewSemester(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateSemester(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.SemesterResponse(), nil
}
