package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type ICourse interface {
	CreateCourse(req *models.CourseRequest) (*models.CourseResponse, error)
	GetCourseById(id uint) (*models.CourseResponse, error)
	ListCourse() ([]models.CourseResponse, error)
	DeleteCourse(id uint) error
	UpdateCourse(id uint, req *models.CourseRequest) (*models.CourseResponse, error)
}

func (s Service) CreateCourse(req *models.CourseRequest) (*models.CourseResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while creating game category")
	}
	newReq, err := models.NewCourse(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateCourse(newReq)
	if err != nil {
		return nil, err
	}
	return data.CourseResponse(), nil
}

func (s Service) GetCourseById(id uint) (*models.CourseResponse, error) {
	data, err := s.repo.FindCourseById(id)
	if err != nil {
		return nil, err
	}
	return data.CourseResponse(), nil
}

func (s Service) ListCourse() ([]models.CourseResponse, error) {
	datas, err := s.repo.FindAllCourse()
	if err != nil {
		return nil, err
	}
	var responses []models.CourseResponse
	for _, data := range *datas {
		responses = append(responses, *data.CourseResponse())
	}
	return responses, nil
}

func (s Service) DeleteCourse(id uint) error {
	return s.repo.DeleteCourse(id)
}

func (s Service) UpdateCourse(id uint, req *models.CourseRequest) (*models.CourseResponse, error) {
	if req == nil {
		return nil, errors.New("game category request is nil while updating game category")
	}
	newReq, err := models.NewCourse(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.UpdateCourse(id, newReq)
	if err != nil {
		return nil, err
	}
	return data.CourseResponse(), nil
}
