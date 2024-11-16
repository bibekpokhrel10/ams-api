package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type ISemester interface {
	CreateSemester(req *models.SemesterRequest) (*models.SemesterResponse, error)
	GetSemesterById(id uint) (*models.SemesterResponse, error)
	ListSemester(req *models.ListSemesterRequest) ([]models.SemesterResponse, int, error)
	DeleteSemester(id uint) error
	UpdateSemester(id uint, req *models.SemesterRequest) (*models.SemesterResponse, error)
}

func (s Service) CreateSemester(req *models.SemesterRequest) (*models.SemesterResponse, error) {
	if req == nil {
		return nil, errors.New("semester request request is nil while creating semester")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repo.FindProgramById(req.ProgramId)
	if err != nil {
		return nil, errors.New("program not found")
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

func (s Service) ListSemester(req *models.ListSemesterRequest) ([]models.SemesterResponse, int, error) {
	if req == nil {
		return nil, 0, errors.New("semester request is nil while listing semester")
	}
	_, err := s.repo.FindProgramById(req.ProgramId)
	if err != nil {
		return nil, 0, errors.New("program not found")
	}
	datas, count, err := s.repo.FindAllSemester(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.SemesterResponse
	for _, data := range *datas {
		responses = append(responses, *data.SemesterResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteSemester(id uint) error {
	return s.repo.DeleteSemester(id)
}

func (s Service) UpdateSemester(id uint, req *models.SemesterRequest) (*models.SemesterResponse, error) {
	if req == nil {
		return nil, errors.New("semester request is nil while updating semester")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindSemesterById(id)
	if err != nil {
		return nil, errors.New("semester not found")
	}
	_, err = s.repo.FindProgramById(req.ProgramId)
	if err != nil {
		return nil, errors.New("program not found")
	}
	data, err := s.repo.UpdateSemester(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.SemesterResponse(), nil
}
