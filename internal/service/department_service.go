package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IDepartment interface {
	CreateDepartment(req *models.DepartmentRequest) (*models.DepartmentResponse, error)
	GetDepartmentById(id uint) (*models.DepartmentResponse, error)
	ListDepartment() ([]models.DepartmentResponse, error)
	DeleteDepartment(id uint) error
	UpdateDepartment(id uint, req *models.DepartmentRequest) (*models.DepartmentResponse, error)
}

func (s Service) CreateDepartment(req *models.DepartmentRequest) (*models.DepartmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewDepartment(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateDepartment(newReq)
	if err != nil {
		return nil, err
	}
	return data.DepartmentResponse(), nil
}

func (s Service) GetDepartmentById(id uint) (*models.DepartmentResponse, error) {
	data, err := s.repo.FindDepartmentById(id)
	if err != nil {
		return nil, err
	}
	return data.DepartmentResponse(), nil
}

func (s Service) ListDepartment() ([]models.DepartmentResponse, error) {
	datas, err := s.repo.FindAllDepartment()
	if err != nil {
		return nil, err
	}
	var responses []models.DepartmentResponse
	for _, data := range *datas {
		responses = append(responses, *data.DepartmentResponse())
	}
	return responses, nil
}

func (s Service) DeleteDepartment(id uint) error {
	return s.repo.DeleteDepartment(id)
}

func (s Service) UpdateDepartment(id uint, req *models.DepartmentRequest) (*models.DepartmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewDepartment(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateDepartment(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.DepartmentResponse(), nil
}
