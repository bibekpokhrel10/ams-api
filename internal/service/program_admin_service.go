package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IProgramAdmin interface {
	CreateProgramAdmin(req *models.ProgramAdminRequest) (*models.ProgramAdminResponse, error)
	GetProgramAdminById(id uint) (*models.ProgramAdminResponse, error)
	ListProgramAdmin(req *models.ListProgramAdminRequest) ([]models.ProgramAdminResponse, int, error)
	DeleteProgramAdmin(id uint) error
	UpdateProgramAdmin(id uint, req *models.ProgramAdminRequest) (*models.ProgramAdminResponse, error)
}

func (s Service) CreateProgramAdmin(req *models.ProgramAdminRequest) (*models.ProgramAdminResponse, error) {
	if req == nil {
		return nil, errors.New("programAdmin request is nil while creating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	newReq, err := models.NewProgramAdmin(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateProgramAdmin(newReq)
	if err != nil {
		return nil, err
	}
	return data.ProgramAdminResponse(), nil
}

func (s Service) GetProgramAdminById(id uint) (*models.ProgramAdminResponse, error) {
	data, err := s.repo.FindProgramAdminById(id)
	if err != nil {
		return nil, err
	}
	return data.ProgramAdminResponse(), nil
}

func (s Service) ListProgramAdmin(req *models.ListProgramAdminRequest) ([]models.ProgramAdminResponse, int, error) {
	if req == nil {
		return nil, 0, errors.New("programAdmin request is nil while listing game category")
	}
	datas, count, err := s.repo.FindAllProgramAdmin(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.ProgramAdminResponse
	for _, data := range *datas {
		responses = append(responses, *data.ProgramAdminResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteProgramAdmin(id uint) error {
	return s.repo.DeleteProgramAdmin(id)
}

func (s Service) UpdateProgramAdmin(id uint, req *models.ProgramAdminRequest) (*models.ProgramAdminResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindProgramAdminById(id)
	if err != nil {
		return nil, errors.New("programAdmin not found")
	}
	data, err := s.repo.UpdateProgramAdmin(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.ProgramAdminResponse(), nil
}
