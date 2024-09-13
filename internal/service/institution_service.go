package service

import (
	"errors"
	"fmt"

	"github.com/ams-api/internal/models"
	"github.com/sirupsen/logrus"
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
		logrus.Error("empty institution request while creating institution")
		return nil, errors.New("institution request is nil while creating institution")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		logrus.Error("error validating institution request :: ", err)
		return nil, err
	}
	newReq := models.NewInstitution(req)
	data, err := s.repo.CreateInstitution(newReq)
	if err != nil {
		logrus.Error("error creating institution :: ", err)
		return nil, err
	}
	return data.InstitutionResponse(), nil
}

func (s Service) GetInstitutionById(id uint) (*models.InstitutionResponse, error) {
	data, err := s.repo.FindInstitutionById(id)
	if err != nil {
		logrus.Errorf("error finding institution by id %d :: %v", id, err)
		return nil, err
	}
	return data.InstitutionResponse(), nil
}

func (s Service) ListInstitution() ([]models.InstitutionResponse, error) {
	datas, err := s.repo.FindAllInstitution()
	if err != nil {
		logrus.Errorf("error finding all institution :: %v", err)
		return nil, err
	}
	var responses []models.InstitutionResponse
	for _, data := range *datas {
		responses = append(responses, *data.InstitutionResponse())
	}
	return responses, nil
}

func (s Service) DeleteInstitution(id uint) error {
	data, err := s.repo.FindInstitutionById(id)
	if err != nil {
		logrus.Errorf("error finding institution by id %d :: %v", id, err)
		return errors.New("institution not found")
	}
	err = s.repo.DeleteInstitution(data.ID)
	if err != nil {
		logrus.Errorf("error deleting institution by id %d :: %v", id, err)
		return err
	}
	return nil
}

func (s Service) UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.InstitutionResponse, error) {
	if req == nil {
		return nil, errors.New("institution request is nil while updating institution")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindInstitutionById(id)
	if err != nil {
		logrus.Errorf("error finding institution by id %d :: %v", id, err)
		return nil, errors.New("institution not found")
	}
	_, err = s.repo.FindInstitutionByName(req.Name)
	if err == nil {
		logrus.Errorf("institution with name %s already exists", req.Name)
		return nil, fmt.Errorf("institution with name %s already exists", req.Name)
	}
	data, err := s.repo.UpdateInstitution(datum.ID, req)
	if err != nil {
		logrus.Errorf("error updating institution by id %d :: %v", id, err)
		return nil, err
	}
	return data.InstitutionResponse(), nil
}
