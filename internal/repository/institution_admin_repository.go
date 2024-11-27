package repository

import (
	"github.com/ams-api/internal/models"
)

type IInstitutionAdmin interface {
	CreateInstitutionAdmin(data *models.InstitutionAdmin) (*models.InstitutionAdmin, error)
	FindAllInstitutionAdmin(req *models.ListInstitutionAdminRequest) (*[]models.InstitutionAdmin, int, error)
	FindInstitutionAdminById(id uint) (*models.InstitutionAdmin, error)
	FindInstitutionAdminByName(name string) (*models.InstitutionAdmin, error)
	UpdateInstitutionAdmin(id uint, req *models.InstitutionAdminRequest) (*models.InstitutionAdmin, error)
	DeleteInstitutionAdmin(id uint) error
}

func (r *Repository) CreateInstitutionAdmin(data *models.InstitutionAdmin) (*models.InstitutionAdmin, error) {
	err := r.db.Model(&models.InstitutionAdmin{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllInstitutionAdmin(req *models.ListInstitutionAdminRequest) (*[]models.InstitutionAdmin, int, error) {
	datas := &[]models.InstitutionAdmin{}
	count := 0
	f := r.db.Model(models.InstitutionAdmin{}).Where("institution_id=?", req.InstitutionId)
	f = f.Count(&count)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindInstitutionAdminById(id uint) (*models.InstitutionAdmin, error) {
	data := &models.InstitutionAdmin{}
	err := r.db.Model(models.InstitutionAdmin{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) FindInstitutionAdminByName(name string) (*models.InstitutionAdmin, error) {
	data := &models.InstitutionAdmin{}
	err := r.db.Model(models.InstitutionAdmin{}).Where("name = ?", name).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateInstitutionAdmin(id uint, req *models.InstitutionAdminRequest) (*models.InstitutionAdmin, error) {
	data := &models.InstitutionAdmin{}
	err := r.db.Model(&models.InstitutionAdmin{}).Where("id = ?", id).Updates(map[string]interface{}{
		"institution_id": req.InstitutionId,
		"user_id":        req.UserId,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteInstitutionAdmin(id uint) error {
	return r.db.Model(&models.InstitutionAdmin{}).Where("id = ?", id).Delete(&models.InstitutionAdmin{}).Error
}
