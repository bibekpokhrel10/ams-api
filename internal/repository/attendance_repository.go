package repository

import (
	"github.com/ams-api/internal/models"
)

type IAttendance interface {
	CreateAttendance(data *models.Attendance) (*models.Attendance, error)
	FindAllAttendance() (*[]models.Attendance, error)
	FindAttendanceById(id uint) (*models.Attendance, error)
	UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.Attendance, error)
	DeleteAttendance(id uint) error
}

func (r *Repository) CreateAttendance(data *models.Attendance) (*models.Attendance, error) {
	err := r.db.Model(&models.Attendance{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllAttendance() (*[]models.Attendance, error) {
	datas := &[]models.Attendance{}
	err := r.db.Model(&models.Attendance{}).Order("id desc").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindAttendanceById(id uint) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(&models.Attendance{}).Where("id = ?", id).Updates(map[string]interface{}{
		"student_id": req.StudentId,
		"class_id":   req.ClassId,
		"date":       req.Date,
		"is_present": req.IsPresent,
		"remarks":    req.Remarks,
	}).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteAttendance(id uint) error {
	return r.db.Model(&models.Attendance{}).Where("id = ?", id).Delete(&models.Attendance{}).Error
}
