package repository

import (
	"strings"

	"github.com/ams-api/internal/models"
)

type IUser interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	CheckEmailAvailableByInstitutionId(institutionId uint, email string) bool
	FindUserByEmailAndInstitutionId(email string, institutionId uint) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	FindAllUsers(req *models.ListUserRequest) ([]models.User, int, error)
	ActivateUser(id uint, isActive bool) (*models.User, error)
	UpdateUserPassword(id uint, password string) (*models.User, error)
	DeleteUser(id uint) error
	UpdateUserType(id uint, userType string) (*models.User, error)
}

func (r *Repository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("email = ?", email).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) CheckEmailAvailableByInstitutionId(institutionId uint, email string) bool {
	var count int64
	if email == "" {
		email = "***"
	}
	r.db.Model(&models.User{}).Where("email = ? AND institution_id = ?", email, institutionId).Count(&count)
	return count > 0
}

func (r *Repository) FindUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	db := r.db.Model(models.User{}).
		Where("username = ? or email = ?", strings.ToLower(username), strings.ToLower(username))
	err := db.Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) FindUserById(id uint) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("id = ?", id).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) FindAllUsers(req *models.ListUserRequest) ([]models.User, int, error) {
	datas := []models.User{}
	count := 0
	f := r.db.Model(models.User{})
	if req.Email != "" {
		f = f.Where("email = ?", req.Email)
	}
	if req.FirstName != "" {
		f = f.Where("first_name = ?", req.FirstName)
	}
	if req.Type != "" {
		f = f.Where("user_type = ?", req.Type)
	}
	f = f.Where("lower(email) LIKE lower(?)", "%"+req.Query+"%").Count(&count)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) ActivateUser(id uint, isActive bool) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("id = ?", id).Update("is_active", isActive).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) UpdateUserPassword(id uint, password string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("id = ?", id).Update("password", password).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) DeleteUser(id uint) error {
	err := r.db.Model(models.User{}).Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return err
}

func (r *Repository) FindUserByEmailAndInstitutionId(email string, institutionId uint) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("email = ? AND institution_id = ?", email, institutionId).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *Repository) UpdateUserType(id uint, userType string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("id = ?", id).Update("user_type", userType).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
