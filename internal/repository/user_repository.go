package repository

import (
	"strings"

	"github.com/ams-api/internal/models"
)

type IUser interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	CheckEmailAvailable(email string) bool
	FindUserByUsername(username string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	FindAllUsers() ([]models.User, error)
	ActivateUser(id uint, isActive bool) (*models.User, error)
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

func (r *Repository) CheckEmailAvailable(email string) bool {
	var count int64
	if email == "" {
		email = "***"
	}
	r.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
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

func (r *Repository) FindAllUsers() ([]models.User, error) {
	datas := []models.User{}
	err := r.db.Model(models.User{}).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) ActivateUser(id uint, isActive bool) (*models.User, error) {
	user := &models.User{}
	err := r.db.Model(models.User{}).Where("id = ?", id).Update("is_active", isActive).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
