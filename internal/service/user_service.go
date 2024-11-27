package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/util"
	"github.com/ams-api/util/constants"
	"github.com/ams-api/util/password"
	"github.com/sirupsen/logrus"
)

type IUser interface {
	CreateUser(req *models.UserRequest) (*models.UserResponse, error)
	LoginUser(req *models.LoginRequest) (*models.LoginResponse, int, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserById(id uint) (*models.UserResponse, error)
	ListAllUsers(req *models.ListUserRequest) ([]models.UserResponse, int, error)
	CheckEmailAvailableByInstitutionId(institutionId uint, email string) bool
	ActivateDeactivateUser(id uint, isActive bool) (*models.UserResponse, error)
	UpdateUserPassword(id uint, req *models.UpdateUserPasswordRequest) (*models.UserResponse, error)
	DeleteUser(id uint) error
	UpdateUserType(userId, id uint, userType string) (*models.UserResponse, error)
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
	_, err := s.repo.FindInstitutionById(req.InstitutionId)
	if err != nil {
		logrus.Error("error finding institution by id :: ", err)
		return nil, err
	}
	if s.repo.CheckEmailAvailableByInstitutionId(req.InstitutionId, req.Email) {
		logrus.Error("email already exists")
		return nil, errors.New("email already exists for this institution")
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

func (s Service) CheckEmailAvailableByInstitutionId(institutionId uint, email string) bool {
	return s.repo.CheckEmailAvailableByInstitutionId(institutionId, email)
}

func (s Service) LoginUser(req *models.LoginRequest) (*models.LoginResponse, int, error) {
	result := &models.LoginResponse{}
	if err := req.Validate(); err != nil {
		logrus.Error("error validating login request :: ", err)
		return nil, http.StatusBadRequest, err
	}
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		logrus.Error("error finding user by email :: ", err)
		return nil, http.StatusNotFound, err
	}
	if req.UserType != user.UserType {
		logrus.Error("user type does not match")
		return nil, http.StatusUnauthorized, fmt.Errorf("user is not %s", req.UserType)
	}
	if req.UserType != constants.SUPER_ADMIN {
		s.repo.FindInstitutionById(req.InsitutionId)
	}
	err = password.CheckPassword(req.Password, user.Password)
	if err != nil {
		logrus.Error("error checking password :: ", err)
		return nil, http.StatusUnauthorized, errors.New("incorrect password")
	}
	if !user.IsActive {
		return nil, http.StatusForbidden, errors.New("please contact to support team to activate your account. ")
	}
	result.User = user.UserResponse()
	return result, http.StatusOK, nil
}

func (s Service) ListAllUsers(req *models.ListUserRequest) ([]models.UserResponse, int, error) {
	datas, count, err := s.repo.FindAllUsers(req)
	if err != nil {
		logrus.Errorf("error finding all users :: %v", err)
		return nil, count, err
	}

	filteredData := []models.User{}
	isFiltered := false

	// Check for program enrollment filter
	if req.IsProgramEnrollment != nil && *req.IsProgramEnrollment {
		programEnrollments, err := s.repo.FindProgramEnrollmentByProgramId(req.ProgramId)
		if err == nil && programEnrollments != nil {
			// Create a map of user IDs from program enrollments for quick lookup
			programEnrolledUsers := []uint{}
			for _, programEnrollment := range programEnrollments {
				programEnrolledUsers = append(programEnrolledUsers, programEnrollment.StudentId)
			}

			// Filter out users in `datas` that are also in `programEnrollments`
			for _, data := range datas {
				isFiltered = true
				if !util.ContainsUint(programEnrolledUsers, data.ID) {
					filteredData = append(filteredData, data)
				}
			}
		}
	}

	// If no filtering occurred, use the original data
	if !isFiltered {
		filteredData = datas
	}

	// Prepare the response
	var responses []models.UserResponse
	for _, data := range filteredData {
		responses = append(responses, *data.UserResponse())
	}

	return responses, count, nil

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

func (s Service) UpdateUserPassword(id uint, req *models.UpdateUserPasswordRequest) (*models.UserResponse, error) {
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	req.Prepare()
	hashpwd, err := password.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}
	req.NewPassword = hashpwd
	data, err := s.repo.UpdateUserPassword(user.ID, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return data.UserResponse(), nil
}

func (s Service) DeleteUser(id uint) error {
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	return s.repo.DeleteUser(id)
}

func (s Service) UpdateUserType(userId, id uint, userType string) (*models.UserResponse, error) {
	userType = strings.ToLower(strings.TrimSpace(userType))
	if err := models.ValidateUserType(userType); err != nil {
		return nil, err
	}
	user, err := s.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	actionUser, err := s.repo.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	if userType == "program_admin" && actionUser.UserType != "institution_admin" {
		return nil, errors.New("only institution admin can create program admin")
	}
	if userType == "institution_admin" && actionUser.UserType != "super_admin" {
		return nil, errors.New("only admin can create institution admin")
	}
	data, err := s.repo.UpdateUserType(user.ID, userType)
	if err != nil {
		return nil, err
	}
	return data.UserResponse(), nil
}
