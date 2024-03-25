package repository

import (
	"github.com/ams-api/internal/models"
)

type IClass interface {
	CreateClass(data *models.Class) (*models.Class, error)
	FindAllClass() (*[]models.Class, error)
	FindClassById(id uint) (*models.Class, error)
	UpdateClass(id uint, data *models.Class) (*models.Class, error)
	DeleteClass(id uint) error
}

func (r *Repository) CreateClass(data *models.Class) (*models.Class, error) {
	err := r.db.Model(&models.Class{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllClass() (*[]models.Class, error) {
	datas := &[]models.Class{}
	err := r.db.Model(&models.Class{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindClassById(id uint) (*models.Class, error) {
	data := &models.Class{}
	err := r.db.Model(models.Class{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateClass(id uint, data *models.Class) (*models.Class, error) {
	err := r.db.Model(&models.Class{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteClass(id uint) error {
	return r.db.Model(&models.Class{}).Where("id = ?", id).Delete(&models.Class{}).Error
}
