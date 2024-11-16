package models

import (
	"errors"
	"fmt"
	"html"
	"regexp"
	"strings"
	"unicode"

	"github.com/ams-api/util/constants"
	"github.com/ams-api/util/password"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserType      string `json:"user_type"` // teacher/student
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"` // male, female, others
	Password      string `json:"password"`
	IsAdmin       bool   `json:"is_admin"`
	IsActive      bool   `json:"is_active"`
	InstitutionId uint   `json:"institution_id"`
}

type UserRequest struct {
	UserType      string `json:"user_type"` // teacher/student
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	Password      string `json:"password"`
	IsAdmin       bool   `json:"is_admin"`
	IsActive      bool   `json:"is_active"`
	InstitutionId uint   `json:"institution_id"`
}

type UserResponse struct {
	Id            uint   `json:"id"`
	UserType      string `json:"user_type"` // teacher/student
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	Password      string `json:"password"`
	IsAdmin       bool   `json:"is_admin"`
	IsActive      bool   `json:"is_active"`
	Role          string `json:"role"`
	InstitutionId uint   `json:"institution_id"`
}

type LoginRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserType     string `json:"user_type"` // super_admin, teacher, student, institution_admin, program_admin
	InsitutionId uint   `json:"institution_id"`
}

type LoginResponse struct {
	Token string        `json:"token"`
	User  *UserResponse `json:"user"`
}
type ActivateDeactivateUserRequest struct {
	IsActive bool `json:"is_active"`
}

func NewUser(req *UserRequest) (*User, error) {
	user := &User{
		UserType:      req.UserType,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		ContactNumber: req.ContactNumber,
		DateOfBirth:   req.DateOfBirth,
		Gender:        req.Gender,
		Password:      req.Password,
		IsAdmin:       req.IsAdmin,
		IsActive:      req.IsActive,
		InstitutionId: req.InstitutionId,
	}
	// user.Prepare()
	// err := user.Validate()
	// if err != nil {
	// 	return nil, err
	// }
	hashpwd, err := password.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashpwd
	return user, nil
}

func (u *User) UserResponse() *UserResponse {
	return &UserResponse{
		Id:            u.ID,
		UserType:      u.UserType,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		ContactNumber: u.ContactNumber,
		DateOfBirth:   u.DateOfBirth,
		Gender:        u.Gender,
		Password:      u.Password,
		IsAdmin:       u.IsAdmin,
		IsActive:      u.IsActive,
	}
}

func (u *UserRequest) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.UserType = strings.TrimSpace(strings.ToLower(u.UserType))
	u.Password = strings.TrimSpace(u.Password)
	// u.IsActive = true
}

func (u *UserRequest) Validate() error {
	if u.Email == "" {
		return ErrEmailRequired
	}
	if err := checkEmail(u.Email); err != nil {
		return ErrInvalidEmail
	}
	if u.Password == "" {
		return ErrPasswordRequired
	}
	if u.UserType == "" {
		return ErrUserTypeRequired
	} else {
		if u.UserType != constants.TEACHER && u.UserType != constants.STUDENT {
			return ErrInvalidUserType
		}
	}
	err := PasswordValidate(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (req *LoginRequest) Prepare() {
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.UserType = strings.TrimSpace(strings.ToLower(req.UserType))
}

func (r *LoginRequest) Validate() error {
	if r.Email == "" {
		return ErrUsernameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	if r.UserType == "" {
		return errors.New("require user type")
	}
	if r.UserType != constants.SUPER_ADMIN && r.UserType != constants.TEACHER && r.UserType != constants.STUDENT && r.UserType != constants.INSTITUTION_ADMIN && r.UserType != constants.PROGRAM_ADMIN {
		return errors.New("invalid user type")
	}
	if r.UserType != constants.SUPER_ADMIN {
		if r.InsitutionId == 0 {
			return errors.New("require insitution id")
		}
	}
	return nil
}

func checkEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

type UpdateUserPasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfrimPassword string `json:"confrim_password"`
}

type ListUserRequest struct {
	ListRequest
	Email     string `form:"email"`
	FirstName string `form:"first_name"`
	Type      string `form:"type"`
}

type UpdateUserType struct {
	UserType string `json:"user_type"`
}

func (req *UpdateUserPasswordRequest) Prepare() {
	req.NewPassword = strings.TrimSpace(req.NewPassword)
	req.ConfirmPassword = strings.TrimSpace(req.ConfirmPassword)
}

func (req *ChangePasswordRequest) Validate() (string, error) {
	if req.OldPassword == "" {
		return "", ErrRequiredOldPassword
	}
	if req.OldPassword == req.NewPassword {
		return "", ErrSamePasswordAsOld
	}
	return validatePassword(req.NewPassword, req.ConfrimPassword)
}

func validatePassword(newPassword, confrimPassword string) (string, error) {
	if newPassword == "" {
		return "", ErrRequiredNewPassword
	}
	if confrimPassword == "" {
		return "", ErrRequiredConfrimPassword
	}
	if newPassword != confrimPassword {
		return "", ErrPasswordMismatch
	}
	err := PasswordValidate(newPassword)
	if err != nil {
		return "", err
	}
	hashPwd, err := password.HashPassword(newPassword)
	if err != nil {
		return "", ErrPasswordHashGenerate
	}
	return hashPwd, nil
}

func PasswordValidate(password string) error {
	if len(password) < 6 {
		return ErrPasswordMinLength
	}
	return nil
}

func PasswordValidateMain(password string) error {
	// Check password length
	if len(password) < 6 {
		return ErrPasswordMinLength
	}

	// Check for at least one uppercase letter
	hasUpper := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		return ErrPasswordUpperCaseLetter
	}

	// Check for at least one lowercase letter
	hasLower := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
			break
		}
	}
	if !hasLower {
		return ErrPasswordLowerCaseLetter
	}

	// Check for at least one digit
	hasDigit := false
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return ErrPasswordOneDigit
	}

	// Check for at least one special character
	hasSpecial := false
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return ErrPasswordSpecialCharacter
	}

	return nil
}

func ValidateUserType(userType string) error {
	if userType != "teacher" && userType != "student" && userType != "institution_admin" && userType != "program_admin" {
		return ErrInvalidUserType
	}
	return nil
}
