package models

import "errors"

var (
	ErrEmailRequired            = errors.New("required email")
	ErrFirstNameRequired        = errors.New("required firstname")
	ErrNameRequired             = errors.New("required name")
	ErrUserTypeRequired         = errors.New("required user type")
	ErrInvalidUserType          = errors.New("invalid user type, must be teacher or student")
	ErrUsernameRequired         = errors.New("required username")
	ErrInvalidEmail             = errors.New("invalid email")
	ErrInvalidUsername          = errors.New("username must be alpha-numeric character only")
	ErrPasswordRequired         = errors.New("required password")
	ErrParentIDRequired         = errors.New("required parent id")
	ErrRequiredOldPassword      = errors.New("required old passowd")
	ErrSamePasswordAsOld        = errors.New("required different than old passowd")
	ErrRequiredNewPassword      = errors.New("required new passowd")
	ErrRequiredConfrimPassword  = errors.New("required Conformation passowd")
	ErrPasswordMinLength        = errors.New("password should be atleast 8 characters")
	ErrPasswordUpperCaseLetter  = errors.New("password must contain at least one uppercase letter")
	ErrPasswordLowerCaseLetter  = errors.New("password must contain at least one lowercase letter")
	ErrPasswordSpecialCharacter = errors.New("password must contain at least one special character")
	ErrPasswordOneDigit         = errors.New("password must contain at least one digit")
	ErrPasswordMismatch         = errors.New("password provided do not match")
	ErrPasswordHashGenerate     = errors.New("unable to generate hash password")
	ErrInvalidEntity            = errors.New("invalid model")

	//Provider Department error
	ErrDepartmentRequired     = errors.New("required department name")
	ErrInvalidDepartmentType  = errors.New("invalid department type, must be teacher or student")
	ErrDepartmentTypeRequired = errors.New("required department type")
)
