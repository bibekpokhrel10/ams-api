package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IClass interface {
	CreateClass(req *models.ClassRequest) (*models.ClassResponse, error)
	GetClassById(id uint) (*models.ClassResponse, error)
	ListClass(req *models.ListClassRequest) ([]models.ClassResponse, error)
	DeleteClass(id uint) error
	UpdateClass(id uint, req *models.ClassRequest) (*models.ClassResponse, error)
}

func (s Service) CreateClass(req *models.ClassRequest) (*models.ClassResponse, error) {
	if req == nil {
		return nil, errors.New("class request is nil while creating class")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repo.FindCourseById(req.CourseId)
	if err != nil {
		return nil, errors.New("course not found")
	}
	user, err := s.repo.FindUserById(req.InstructorID)
	if err != nil {
		return nil, errors.New("instructor not found")
	}
	if user.UserType != "teacher" {
		return nil, errors.New("instructor not found")
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

func (s Service) ListClass(req *models.ListClassRequest) ([]models.ClassResponse, error) {
	datas, err := s.repo.FindAllClass(req)
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
		return nil, errors.New("class request is nil while updating class")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindClassById(id)
	if err != nil {
		return nil, errors.New("class not found")
	}
	_, err = s.repo.FindCourseById(req.CourseId)
	if err != nil {
		return nil, errors.New("course not found")
	}
	user, err := s.repo.FindUserById(req.InstructorID)
	if err != nil {
		return nil, errors.New("instructor not found")
	}
	if user.UserType != "teacher" {
		return nil, errors.New("instructor not found")
	}
	data, err := s.repo.UpdateClass(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.ClassResponse(), nil
}
