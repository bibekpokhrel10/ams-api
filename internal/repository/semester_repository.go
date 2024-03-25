package repository

import (
	"github.com/ams-api/internal/models"
)

type ISemester interface {
	CreateSemester(data *models.Semester) (*models.Semester, error)
	FindAllSemester() (*[]models.Semester, error)
	FindSemesterById(id uint) (*models.Semester, error)
	UpdateSemester(id uint, data *models.Semester) (*models.Semester, error)
	DeleteSemester(id uint) error
}

func (r *Repository) CreateSemester(data *models.Semester) (*models.Semester, error) {
	err := r.db.Model(&models.Semester{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllSemester() (*[]models.Semester, error) {
	datas := &[]models.Semester{}
	err := r.db.Model(&models.Semester{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindSemesterById(id uint) (*models.Semester, error) {
	data := &models.Semester{}
	err := r.db.Model(models.Semester{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateSemester(id uint, data *models.Semester) (*models.Semester, error) {
	err := r.db.Model(&models.Semester{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteSemester(id uint) error {
	return r.db.Model(&models.Semester{}).Where("id = ?", id).Delete(&models.Semester{}).Error
}
