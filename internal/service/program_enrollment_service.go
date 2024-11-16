package service

import (
	"errors"
	"fmt"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/util/constants"
)

type IProgramEnrollment interface {
	CreateProgramEnrollment(req *models.ProgramEnrollmentRequestPayload) error
	GetProgramEnrollmentById(id uint) (*models.ProgramEnrollmentResponse, error)
	ListProgramEnrollment(req *models.ListProgramEnrollmentRequest) ([]models.ProgramEnrollmentResponse, int, error)
	DeleteProgramEnrollment(id uint) error
	UpdateProgramEnrollment(id uint, req *models.ProgramEnrollmentRequest) (*models.ProgramEnrollmentResponse, error)
}

func (s Service) CreateProgramEnrollment(req *models.ProgramEnrollmentRequestPayload) error {
	if req == nil {
		return errors.New("program enrollment request is nil while creating Program Enrollment")
	}
	_, err := s.repo.FindProgramById(req.ProgramId)
	if err != nil {
		return errors.New("program not found")
	}
	for _, studentId := range req.StudentIds {
		user, err := s.repo.FindUserById(studentId)
		if err != nil {
			return fmt.Errorf("user not found :: %w", err)
		}
		if user.UserType != constants.STUDENT {
			return fmt.Errorf("%s user is not student", user.UserType)
		}
		_, err = s.repo.FindProgramEnrollmentByProgramIdAndStudentId(req.ProgramId, studentId)
		if err == nil {
			return fmt.Errorf("student is already enrolled in this program")
		}
	}
	for _, studentId := range req.StudentIds {
		newReq, err := models.NewProgramEnrollment(&models.ProgramEnrollmentRequest{
			ProgramId: req.ProgramId,
			StudentId: studentId,
		})
		if err != nil {
			return err
		}
		_, err = s.repo.CreateProgramEnrollment(newReq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) GetProgramEnrollmentById(id uint) (*models.ProgramEnrollmentResponse, error) {
	data, err := s.repo.FindProgramEnrollmentById(id)
	if err != nil {
		return nil, err
	}
	return data.ProgramEnrollmentResponse(), nil
}

func (s Service) ListProgramEnrollment(req *models.ListProgramEnrollmentRequest) ([]models.ProgramEnrollmentResponse, int, error) {
	datas, count, err := s.repo.FindAllProgramEnrollment(req)
	if err != nil {
		return nil, count, err
	}
	var responses []models.ProgramEnrollmentResponse
	for _, data := range datas {
		responses = append(responses, *data.ProgramEnrollmentResponse())
	}
	return responses, count, nil
}

func (s Service) DeleteProgramEnrollment(id uint) error {
	return s.repo.DeleteProgramEnrollment(id)
}

func (s Service) UpdateProgramEnrollment(id uint, req *models.ProgramEnrollmentRequest) (*models.ProgramEnrollmentResponse, error) {
	if req == nil {
		return nil, errors.New("program enrollment request is nil while updating ProgramEnrollment")
	}
	datum, err := s.repo.FindProgramEnrollmentById(id)
	if err != nil {
		return nil, errors.New("program enrollment not found")
	}
	_, err = s.repo.FindProgramById(req.ProgramId)
	if err != nil {
		return nil, errors.New("program not found")
	}
	user, err := s.repo.FindUserById(req.StudentId)
	if err != nil {
		return nil, errors.New("student not found")
	}
	if user.UserType != constants.STUDENT {
		return nil, errors.New("student not found")
	}
	data, err := s.repo.UpdateProgramEnrollment(datum.ID, req)
	if err != nil {
		return nil, err
	}
	return data.ProgramEnrollmentResponse(), nil
}
