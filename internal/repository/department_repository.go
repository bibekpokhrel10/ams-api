package repository

import (
	"github.com/ams-api/internal/models"
)

type IDepartment interface {
	CreateDepartment(data *models.Department) (*models.Department, error)
	FindAllDepartment() (*[]models.Department, error)
	FindDepartmentByID(id uint) (*models.Department, error)
	UpdateDepartment(id uint, data *models.Department) (*models.Department, error)
	DeleteDepartment(id uint) error
}

func (r *Repository) CreateDepartment(data *models.Department) (*models.Department, error) {
	err := r.db.Model(&models.Department{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllDepartment() (*[]models.Department, error) {
	datas := &[]models.Department{}
	err := r.db.Model(&models.Department{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindDepartmentByID(id uint) (*models.Department, error) {
	data := &models.Department{}
	err := r.db.Model(models.Department{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateDepartment(id uint, data *models.Department) (*models.Department, error) {
	err := r.db.Model(&models.Department{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteDepartment(id uint) error {
	return r.db.Model(&models.Department{}).Where("id = ?", id).Delete(&models.Department{}).Error
}
