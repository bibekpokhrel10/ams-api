package repository

import (
	"github.com/ams-api/internal/models"
)

type ISemester interface {
	CreateSemester(data *models.Semester) (*models.Semester, error)
	FindAllSemester(req *models.ListSemesterRequest) (*[]models.Semester, int, error)
	FindSemesterById(id uint) (*models.Semester, error)
	UpdateSemester(id uint, req *models.SemesterRequest) (*models.Semester, error)
	DeleteSemester(id uint) error
}

func (r *Repository) CreateSemester(data *models.Semester) (*models.Semester, error) {
	err := r.db.Model(&models.Semester{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllSemester(req *models.ListSemesterRequest) (*[]models.Semester, int, error) {
	datas := &[]models.Semester{}
	count := 0
	f := r.db.Model(&models.Semester{}).Where("program_id = ?", req.ProgramId)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindSemesterById(id uint) (*models.Semester, error) {
	data := &models.Semester{}
	err := r.db.Model(models.Semester{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateSemester(id uint, req *models.SemesterRequest) (*models.Semester, error) {
	data := &models.Semester{}
	err := r.db.Model(&models.Semester{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        req.Name,
		"time_period": req.TimePeriod,
		"program_id":  req.ProgramId,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteSemester(id uint) error {
	return r.db.Model(&models.Semester{}).Where("id = ?", id).Delete(&models.Semester{}).Error
}
