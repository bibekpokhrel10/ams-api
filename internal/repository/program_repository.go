package repository

import (
	"github.com/ams-api/internal/models"
)

type IProgram interface {
	CreateProgram(data *models.Program) (*models.Program, error)
	FindAllProgram(req *models.ListProgramRequest) (*[]models.Program, int, error)
	FindProgramById(id uint) (*models.Program, error)
	UpdateProgram(id uint, req *models.ProgramRequest) (*models.Program, error)
	DeleteProgram(id uint) error
}

func (r *Repository) CreateProgram(data *models.Program) (*models.Program, error) {
	err := r.db.Model(&models.Program{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllProgram(req *models.ListProgramRequest) (*[]models.Program, int, error) {
	datas := &[]models.Program{}
	count := 0
	f := r.db.Model(models.Program{}).Where("institution_id=?", req.InstitutionId)
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

func (r *Repository) FindProgramById(id uint) (*models.Program, error) {
	data := &models.Program{}
	err := r.db.Model(models.Program{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateProgram(id uint, req *models.ProgramRequest) (*models.Program, error) {
	data := &models.Program{}
	err := r.db.Model(&models.Program{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name": req.Name,
		"type": req.Type,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteProgram(id uint) error {
	return r.db.Model(&models.Program{}).Where("id = ?", id).Delete(&models.Program{}).Error
}
