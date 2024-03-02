package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IDepartment interface {
	CreateDepartment(req *models.DepartmentRequest) (*models.DepartmentResponse, error)
	GetDepartmentById(gcid uint) (*models.DepartmentResponse, error)
	ListDepartment() ([]models.DepartmentResponse, error)
	DeleteDepartment(gcid uint) error
	UpdateDepartment(gcid uint, req *models.DepartmentRequest) (*models.DepartmentResponse, error)
}

func (s Service) CreateDepartment(req *models.DepartmentRequest) (*models.DepartmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	data, err := s.repo.CreateDepartment(models.NewDepartment(req))
	if err != nil {
		return nil, err
	}
	return data.DepartmentResponse(), nil
}

func (s Service) GetDepartmentById(gcid uint) (*models.DepartmentResponse, error) {
	data, err := s.repo.FindDepartmentByID(gcid)
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
	var DepartmentResponses []models.DepartmentResponse
	for _, data := range *datas {
		DepartmentResponses = append(DepartmentResponses, *data.DepartmentResponse())
	}
	return DepartmentResponses, nil
}

func (s Service) DeleteDepartment(gcid uint) error {
	return s.repo.DeleteDepartment(gcid)
}

func (s Service) UpdateDepartment(gcid uint, req *models.DepartmentRequest) (*models.DepartmentResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	data, err := s.repo.UpdateDepartment(gcid, models.NewDepartment(req))
	if err != nil {
		return nil, err
	}
	return data.DepartmentResponse(), nil
}
