package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IClass interface {
	CreateClass(req *models.ClassRequest) (*models.ClassResponse, error)
	GetClassById(id uint) (*models.ClassResponse, error)
	ListClass() ([]models.ClassResponse, error)
	DeleteClass(id uint) error
	UpdateClass(id uint, req *models.ClassRequest) (*models.ClassResponse, error)
}

func (s Service) CreateClass(req *models.ClassRequest) (*models.ClassResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewClass(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateClass(newReq)
	if err != nil {
		return nil, err
	}
	return data.ClassResponse(), nil
}

func (s Service) GetClassById(id uint) (*models.ClassResponse, error) {
	data, err := s.repo.FindClassById(id)
	if err != nil {
		return nil, err
	}
	return data.ClassResponse(), nil
}

func (s Service) ListClass() ([]models.ClassResponse, error) {
	datas, err := s.repo.FindAllClass()
	if err != nil {
		return nil, err
	}
	var responses []models.ClassResponse
	for _, data := range *datas {
		responses = append(responses, *data.ClassResponse())
	}
	return responses, nil
}

func (s Service) DeleteClass(id uint) error {
	return s.repo.DeleteClass(id)
}

func (s Service) UpdateClass(id uint, req *models.ClassRequest) (*models.ClassResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewClass(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateClass(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.ClassResponse(), nil
}
