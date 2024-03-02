package service

import (
	"errors"
	"net/http"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/util/password"
)

type IUser interface {
	CreateUser(req *models.UserRequest) (*models.UserResponse, error)
	LoginUser(req *models.LoginRequest) (*models.LoginResponse, int, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	CheckEmailAvailable(email string) bool
}

func (s Service) CreateUser(req *models.UserRequest) (*models.UserResponse, error) {
	user, err := models.NewUser(req)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data.UserResponse(), nil
}

func (s Service) GetUserByEmail(email string) (*models.UserResponse, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user.UserResponse(), nil
}

func (s Service) CheckEmailAvailable(email string) bool {
	return s.repo.CheckEmailAvailable(email)
}

func (s Service) LoginUser(req *models.LoginRequest) (*models.LoginResponse, int, error) {
	result := &models.LoginResponse{}
	if err := req.Validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	err = password.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("incorrect password")
	}
	if !user.IsActive {
		return nil, http.StatusForbidden, errors.New("please contact to support team to activate your account. ")
	}
	result.User = user.UserResponse()
	return result, http.StatusOK, nil
}
