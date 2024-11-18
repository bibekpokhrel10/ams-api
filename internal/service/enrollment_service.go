package service

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IEnrollment interface {
	CreateEnrollment(req *models.EnrollmentRequestPayload) error
	GetEnrollmentById(id uint) (*models.EnrollmentResponse, error)
	ListEnrollment(req *models.ListEnrollmentRequest) ([]models.EnrollmentResponse, int, error)
	DeleteEnrollment(id uint) error
	UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.EnrollmentResponse, error)
}

func (s Service) CreateEnrollment(req *models.EnrollmentRequestPayload) error {
	if req == nil {
		return errors.New("enrollment request is nil while creating enrollment")
	}
	_, err := s.repo.FindClassById(req.ClassId)
	if err != nil {
		return errors.New("class not found")
	}
	for _, studentId := range req.StudentIds {
		user, err := s.repo.FindUserById(studentId)
		if err != nil {
			return errors.New("student not found")
		}
		if user.UserType != "student" {
			return errors.New("student not found")
		}
	}
	for _, studentId := range req.StudentIds {
		newReq, err := models.NewEnrollment(&models.EnrollmentRequest{
			ClassId:   req.ClassId,
			StudentId: studentId,
		})
		if err != nil {
			return err
		}
		_, err = s.repo.CreateEnrollment(newReq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetEnrollmentById(id uint) (*models.EnrollmentResponse, error) {
	data, err := s.repo.FindEnrollmentById(id)
	if err != nil {
		return nil, err
	}
	return data.EnrollmentResponse(), nil
}

func (s Service) ListEnrollment(req *models.ListEnrollmentRequest) ([]models.EnrollmentResponse, int, error) {
	datas, count, err := s.repo.FindAllEnrollment(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.EnrollmentResponse
	for _, data := range datas {
		responses = append(responses, *data.EnrollmentResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteEnrollment(id uint) error {
	return s.repo.DeleteEnrollment(id)
}

func (s Service) UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.EnrollmentResponse, error) {
	if req == nil {
		return nil, errors.New("enrollment request is nil while updating enrollment")
	}
	datum, err := s.repo.FindEnrollmentById(id)
	if err != nil {
		return nil, errors.New("enrollment not found")
	}
	_, err = s.repo.FindClassById(req.ClassId)
	if err != nil {
		return nil, errors.New("class not found")
	}
	user, err := s.repo.FindUserById(req.StudentId)
	if err != nil {
		return nil, errors.New("student not found")
	}
	if user.UserType != "student" {
		return nil, errors.New("student not found")
	}
	data, err := s.repo.UpdateEnrollment(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.EnrollmentResponse(), nil
}
