package repository

import (
	"github.com/ams-api/internal/models"
)

type IEnrollment interface {
	CreateEnrollment(data *models.Enrollment) (*models.Enrollment, error)
	FindAllEnrollment() (*[]models.Enrollment, error)
	FindEnrollmentById(id uint) (*models.Enrollment, error)
	UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.Enrollment, error)
	DeleteEnrollment(id uint) error
}

func (r *Repository) CreateEnrollment(data *models.Enrollment) (*models.Enrollment, error) {
	err := r.db.Model(&models.Enrollment{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllEnrollment() (*[]models.Enrollment, error) {
	datas := &[]models.Enrollment{}
	err := r.db.Model(&models.Enrollment{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindEnrollmentById(id uint) (*models.Enrollment, error) {
	data := &models.Enrollment{}
	err := r.db.Model(models.Enrollment{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.Enrollment, error) {
	data := &models.Enrollment{}
	err := r.db.Model(&models.Enrollment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"student_id": req.StudentId,
		"class_id":   req.ClassId,
	}).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteEnrollment(id uint) error {
	return r.db.Model(&models.Enrollment{}).Where("id = ?", id).Delete(&models.Enrollment{}).Error
}
