package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IInstitution interface {
	CreateInstitution(req *models.InstitutionRequest) (*models.InstitutionResponse, error)
	GetInstitutionById(id uint) (*models.InstitutionResponse, error)
	ListInstitution() ([]models.InstitutionResponse, error)
	DeleteInstitution(id uint) error
	UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.InstitutionResponse, error)
}

func (s Service) CreateInstitution(req *models.InstitutionRequest) (*models.InstitutionResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	newReq, err := models.NewInstitution(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateInstitution(newReq)
	if err != nil {
		return nil, err
	}
	return data.InstitutionResponse(), nil
}

func (s Service) GetInstitutionById(id uint) (*models.InstitutionResponse, error) {
	data, err := s.repo.FindInstitutionById(id)
	if err != nil {
		return nil, err
	}
	return data.InstitutionResponse(), nil
}

func (s Service) ListInstitution() ([]models.InstitutionResponse, error) {
	datas, err := s.repo.FindAllInstitution()
	if err != nil {
		return nil, err
	}
	var responses []models.InstitutionResponse
	for _, data := range *datas {
		responses = append(responses, *data.InstitutionResponse())
	}
	return responses, nil
}

func (s Service) DeleteInstitution(id uint) error {
	return s.repo.DeleteInstitution(id)
}

func (s Service) UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.InstitutionResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindInstitutionById(id)
	if err != nil {
		return nil, errors.New("institution not found")
	}
	data, err := s.repo.UpdateInstitution(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.InstitutionResponse(), nil
}
