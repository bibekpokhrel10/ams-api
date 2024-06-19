package repository

import (
	"github.com/ams-api/internal/models"
)

type IProgram interface {
	CreateProgram(data *models.Program) (*models.Program, error)
	FindAllProgram() (*[]models.Program, error)
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

func (r *Repository) FindAllProgram() (*[]models.Program, error) {
	datas := &[]models.Program{}
	err := r.db.Model(&models.Program{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
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
