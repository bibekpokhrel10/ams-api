package service

import (
	"errors"
	"net/http"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/util/password"
	"github.com/sirupsen/logrus"
)

type IUser interface {
	CreateUser(req *models.UserRequest) (*models.UserResponse, error)
	LoginUser(req *models.LoginRequest) (*models.LoginResponse, int, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserById(id uint) (*models.UserResponse, error)
	ListAllUsers() ([]models.UserResponse, error)
	CheckEmailAvailable(email string) bool
	ActivateDeactivateUser(id uint, isActive bool) (*models.UserResponse, error)
}

func (s Service) CreateUser(req *models.UserRequest) (*models.UserResponse, error) {
	if req == nil {
		return nil, errors.New("empty user request")
	}
	req.Prepare()
	if err := req.Validate(); err != nil {
		logrus.Error("error validating user request :: ", err)
		return nil, err
	}
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

func (s Service) GetUserById(id uint) (*models.UserResponse, error) {
	user, err := s.repo.FindUserById(id)
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

func (s Service) ListAllUsers() ([]models.UserResponse, error) {
	datas, err := s.repo.FindAllUsers()
	if err != nil {
		logrus.Errorf("error finding all users :: %v", err)
		return nil, err
	}
	var responses []models.UserResponse
	for _, data := range datas {
		responses = append(responses, *data.UserResponse())
	}
	return responses, nil
}

func (s Service) ActivateDeactivateUser(id uint, isActive bool) (*models.UserResponse, error) {
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	data, err := s.repo.ActivateUser(user.ID, isActive)
	if err != nil {
		return nil, err
	}
	return data.UserResponse(), nil
}
