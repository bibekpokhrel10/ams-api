package service

import (
	"errors"

	"github.com/ams-api/internal/models"
	"github.com/sirupsen/logrus"
)

type IProgram interface {
	CreateProgram(req *models.ProgramRequest) (*models.ProgramResponse, error)
	GetProgramById(id uint) (*models.ProgramResponse, error)
	ListProgram(req *models.ListProgramRequest) ([]models.ProgramResponse, int, error)
	DeleteProgram(id uint) error
	UpdateProgram(id uint, req *models.ProgramRequest) (*models.ProgramResponse, error)
}

func (s Service) CreateProgram(req *models.ProgramRequest) (*models.ProgramResponse, error) {
	if req == nil {
		return nil, errors.New("program request is nil while creating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repo.FindInstitutionById(req.InstitutionId)
	if err != nil {
		return nil, err
	}
	newReq, err := models.NewProgram(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateProgram(newReq)
	if err != nil {
		return nil, err
	}
	return data.ProgramResponse(), nil
}

func (s Service) GetProgramById(id uint) (*models.ProgramResponse, error) {
	data, err := s.repo.FindProgramById(id)
	if err != nil {
		return nil, err
	}
	return data.ProgramResponse(), nil
}

func (s Service) ListProgram(req *models.ListProgramRequest) ([]models.ProgramResponse, int, error) {
	if req == nil {
		return nil, 0, errors.New("program request is nil while listing game category")
	}
	_, err := s.repo.FindInstitutionById(req.InstitutionId)
	if err != nil {
		return nil, 0, errors.New("institution not found")
	}
	datas, count, err := s.repo.FindAllProgram(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.ProgramResponse
	for _, data := range *datas {
		responses = append(responses, *data.ProgramResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteProgram(id uint) error {
	return s.repo.DeleteProgram(id)
}

func (s Service) UpdateProgram(id uint, req *models.ProgramRequest) (*models.ProgramResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	logrus.Info("program id :: ", id)
	logrus.Info("program name :: ", req.Name)
	datum, err := s.repo.FindProgramById(id)
	if err != nil {
		return nil, errors.New("program not found")
	}
	data, err := s.repo.UpdateProgram(datum.ID, req)
	if err != nil {
		return nil, err
	}
	logrus.Info("program name :: ", data.Name)
	return data.ProgramResponse(), nil
}
