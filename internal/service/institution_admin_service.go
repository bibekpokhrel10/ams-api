package service

import (
	"errors"

	"github.com/ams-api/internal/models"
	"github.com/sirupsen/logrus"
)

type IInstitutionAdmin interface {
	CreateInstitutionAdmin(req *models.InstitutionAdminRequest) (*models.InstitutionAdminResponse, error)
	GetInstitutionAdminById(id uint) (*models.InstitutionAdminResponse, error)
	ListInstitutionAdmin(req *models.ListInstitutionAdminRequest) ([]models.InstitutionAdminResponse, int, error)
	DeleteInstitutionAdmin(id uint) error
	UpdateInstitutionAdmin(id uint, req *models.InstitutionAdminRequest) (*models.InstitutionAdminResponse, error)
}

func (s Service) CreateInstitutionAdmin(req *models.InstitutionAdminRequest) (*models.InstitutionAdminResponse, error) {
	if req == nil {
		logrus.Error("empty InstitutionAdmin request while creating InstitutionAdmin")
		return nil, errors.New("InstitutionAdmin request is nil while creating InstitutionAdmin")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		logrus.Error("error validating InstitutionAdmin request :: ", err)
		return nil, err
	}
	newReq := models.NewInstitutionAdmin(req)
	data, err := s.repo.CreateInstitutionAdmin(newReq)
	if err != nil {
		logrus.Error("error creating InstitutionAdmin :: ", err)
		return nil, err
	}
	return data.InstitutionAdminResponse(), nil
}

func (s Service) GetInstitutionAdminById(id uint) (*models.InstitutionAdminResponse, error) {
	data, err := s.repo.FindInstitutionAdminById(id)
	if err != nil {
		logrus.Errorf("error finding InstitutionAdmin by id %d :: %v", id, err)
		return nil, err
	}
	return data.InstitutionAdminResponse(), nil
}

func (s Service) ListInstitutionAdmin(req *models.ListInstitutionAdminRequest) ([]models.InstitutionAdminResponse, int, error) {
	datas, count, err := s.repo.FindAllInstitutionAdmin(req)
	if err != nil {
		logrus.Errorf("error finding all InstitutionAdmin :: %v", err)
		return nil, count, err
	}
	var responses []models.InstitutionAdminResponse
	for _, data := range *datas {
		responses = append(responses, *data.InstitutionAdminResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteInstitutionAdmin(id uint) error {
	data, err := s.repo.FindInstitutionAdminById(id)
	if err != nil {
		logrus.Errorf("error finding InstitutionAdmin by id %d :: %v", id, err)
		return errors.New("InstitutionAdmin not found")
	}
	err = s.repo.DeleteInstitutionAdmin(data.ID)
	if err != nil {
		logrus.Errorf("error deleting InstitutionAdmin by id %d :: %v", id, err)
		return err
	}
	return nil
}

func (s Service) UpdateInstitutionAdmin(id uint, req *models.InstitutionAdminRequest) (*models.InstitutionAdminResponse, error) {
	if req == nil {
		return nil, errors.New("InstitutionAdmin request is nil while updating InstitutionAdmin")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindInstitutionAdminById(id)
	if err != nil {
		logrus.Errorf("error finding InstitutionAdmin by id %d :: %v", id, err)
		return nil, errors.New("InstitutionAdmin not found")
	}
	data, err := s.repo.UpdateInstitutionAdmin(datum.ID, req)
	if err != nil {
		logrus.Errorf("error updating InstitutionAdmin by id %d :: %v", id, err)
		return nil, err
	}
	return data.InstitutionAdminResponse(), nil
}
