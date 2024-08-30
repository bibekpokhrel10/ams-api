package repository

import (
	"github.com/ams-api/internal/models"
)

type IInstitution interface {
	CreateInstitution(data *models.Institution) (*models.Institution, error)
	FindAllInstitution() (*[]models.Institution, error)
	FindInstitutionById(id uint) (*models.Institution, error)
	UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.Institution, error)
	DeleteInstitution(id uint) error
}

func (r *Repository) CreateInstitution(data *models.Institution) (*models.Institution, error) {
	err := r.db.Model(&models.Institution{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllInstitution() (*[]models.Institution, error) {
	datas := &[]models.Institution{}
	err := r.db.Model(&models.Institution{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindInstitutionById(id uint) (*models.Institution, error) {
	data := &models.Institution{}
	err := r.db.Model(models.Institution{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.Institution, error) {
	data := &models.Institution{}
	err := r.db.Model(&models.Institution{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name": req.Name,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteInstitution(id uint) error {
	return r.db.Model(&models.Institution{}).Where("id = ?", id).Delete(&models.Institution{}).Error
}
