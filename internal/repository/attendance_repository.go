package repository

import (
	"time"

	"github.com/ams-api/internal/models"
)

type IAttendance interface {
	CreateAttendance(data *models.Attendance) (*models.Attendance, error)
	FindAllAttendance(req *models.ListAttendanceRequest) (*[]models.Attendance, int, error)
	FindAllAttendancesByClassIdAndDate(classId uint, date string) (*[]models.Attendance, error)
	FindAttendancesByClassIdAndDateAndStudentId(classId uint, date time.Time, studentId uint) (*models.Attendance, error)
	FindAttendanceById(id uint) (*models.Attendance, error)
	UpdateAttendance(id uint, req *models.AttendanceRequest) (*models.Attendance, error)
	DeleteAttendance(id uint) error
}

func (r *Repository) CreateAttendance(data *models.Attendance) (*models.Attendance, error) {
	err := r.db.Model(&models.Attendance{}).Save(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAttendancesByClassIdAndDateAndStudentId(classId uint, date time.Time, studentId uint) (*models.Attendance, error) {
	data := &models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("class_id = ? AND date = ? AND student_id = ?", classId, date, studentId).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err

}

func (r *Repository) FindAllAttendance(req *models.ListAttendanceRequest) (*[]models.Attendance, int, error) {
	datas := &[]models.Attendance{}
	count := 0
	f := r.db.Model(models.Attendance{})
	if req.ClassId != 0 {
		f = f.Where("class_id = ?", req.ClassId)
	}
	if req.StudentId != 0 {
		f = f.Where("student_id = ?", req.StudentId)
	}
	if req.Date != "" {
		f = f.Where("date = ?", req.Date)
	}
	f = f.Count(&count)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindAllAttendancesByClassIdAndDate(classId uint, date string) (*[]models.Attendance, error) {
	datas := &[]models.Attendance{}
	err := r.db.Model(models.Attendance{}).Where("class_id = ? AND date = ?", classId, date).Find(&datas).Error
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
		"status":     req.Status,
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
