package repository

import (
	"errors"

	"github.com/ams-api/internal/models"
)

type IInstitution interface {
	CreateInstitution(data *models.Institution) (*models.Institution, error)
	FindAllInstitution(req *models.ListInstitutionRequest) (*[]models.Institution, int, error)
	FindInstitutionById(id uint) (*models.Institution, error)
	FindInstitutionByName(name string) (*models.Institution, error)
	UpdateInstitution(id uint, req *models.InstitutionRequest) (*models.Institution, error)
	DeleteInstitution(id uint) error
}

func (r *Repository) CreateInstitution(data *models.Institution) (*models.Institution, error) {
	err := r.db.Model(&models.Institution{}).Where("name = ?", data.Name).Take(&data).Error
	if err == nil {
		return nil, errors.New("institution already exist")
	}
	err = r.db.Model(&models.Institution{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllInstitution(req *models.ListInstitutionRequest) (*[]models.Institution, int, error) {
	datas := &[]models.Institution{}
	count := 0
	f := r.db.Model(&models.Institution{})
	if req.Name != "" {
		f = f.Where("name=?", req.Name)
	}
	f = f.Where("lower(name) LIKE lower(?)", "%"+req.Query+"%").Count(&count)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindInstitutionById(id uint) (*models.Institution, error) {
	data := &models.Institution{}
	err := r.db.Model(models.Institution{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) FindInstitutionByName(name string) (*models.Institution, error) {
	data := &models.Institution{}
	err := r.db.Model(models.Institution{}).Where("name = ?", name).Take(data).Error
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
