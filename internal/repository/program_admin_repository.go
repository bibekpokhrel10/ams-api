package repository

import (
	"github.com/ams-api/internal/models"
)

type IProgramAdmin interface {
	CreateProgramAdmin(data *models.ProgramAdmin) (*models.ProgramAdmin, error)
	FindAllProgramAdmin(req *models.ListProgramAdminRequest) (*[]models.ProgramAdmin, int, error)
	FindProgramAdminById(id uint) (*models.ProgramAdmin, error)
	UpdateProgramAdmin(id uint, req *models.ProgramAdminRequest) (*models.ProgramAdmin, error)
	DeleteProgramAdmin(id uint) error
}

func (r *Repository) CreateProgramAdmin(data *models.ProgramAdmin) (*models.ProgramAdmin, error) {
	err := r.db.Model(&models.ProgramAdmin{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllProgramAdmin(req *models.ListProgramAdminRequest) (*[]models.ProgramAdmin, int, error) {
	datas := &[]models.ProgramAdmin{}
	count := 0
	f := r.db.Model(models.ProgramAdmin{}).Where("program_id=?", req.ProgramId)
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

func (r *Repository) FindProgramAdminById(id uint) (*models.ProgramAdmin, error) {
	data := &models.ProgramAdmin{}
	err := r.db.Model(models.ProgramAdmin{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateProgramAdmin(id uint, req *models.ProgramAdminRequest) (*models.ProgramAdmin, error) {
	data := &models.ProgramAdmin{}
	err := r.db.Model(&models.ProgramAdmin{}).Where("id = ?", id).Updates(map[string]interface{}{
		"program_id": req.ProgramId,
		"user_id":    req.UserId,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteProgramAdmin(id uint) error {
	return r.db.Model(&models.ProgramAdmin{}).Where("id = ?", id).Delete(&models.ProgramAdmin{}).Error
}
