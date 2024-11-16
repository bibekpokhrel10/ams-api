package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type ICourse interface {
	CreateCourse(req *models.CourseRequest) (*models.CourseResponse, error)
	GetCourseById(id uint) (*models.CourseResponse, error)
	ListCourse(req *models.ListCourseRequest) ([]models.CourseResponse, int, error)
	DeleteCourse(id uint) error
	UpdateCourse(id uint, req *models.CourseRequest) (*models.CourseResponse, error)
}

func (s Service) CreateCourse(req *models.CourseRequest) (*models.CourseResponse, error) {
	if req == nil {
		return nil, errors.New("course request is nil while creating course")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	_, err := s.repo.FindSemesterById(req.SemesterId)
	if err != nil {
		return nil, errors.New("semester not found")
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

func (s Service) ListCourse(req *models.ListCourseRequest) ([]models.CourseResponse, int, error) {
	datas, count, err := s.repo.FindAllCourse(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.CourseResponse
	for _, data := range datas {
		responses = append(responses, *data.CourseResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteCourse(id uint) error {
	return s.repo.DeleteCourse(id)
}

func (s Service) UpdateCourse(id uint, req *models.CourseRequest) (*models.CourseResponse, error) {
	if req == nil {
		return nil, errors.New("course request is nil while updating course")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	datum, err := s.repo.FindCourseById(id)
	if err != nil {
		return nil, errors.New("course not found")
	}
	_, err = s.repo.FindSemesterById(req.SemesterId)
	if err != nil {
		return nil, errors.New("semester not found")
	}
	data, err := s.repo.UpdateCourse(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.CourseResponse(), nil
}
