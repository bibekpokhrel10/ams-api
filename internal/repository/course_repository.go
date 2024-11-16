package repository

import (
	"github.com/ams-api/internal/models"
)

type ICourse interface {
	CreateCourse(data *models.Course) (*models.Course, error)
	FindAllCourse(req *models.ListCourseRequest) ([]models.Course, int, error)
	FindCourseById(id uint) (*models.Course, error)
	UpdateCourse(id uint, req *models.CourseRequest) (*models.Course, error)
	DeleteCourse(id uint) error
}

func (r *Repository) CreateCourse(data *models.Course) (*models.Course, error) {
	err := r.db.Model(&models.Course{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllCourse(req *models.ListCourseRequest) ([]models.Course, int, error) {
	datas := []models.Course{}
	count := 0
	f := r.db.Model(&models.Course{}).Where("semester_id=?", req.SemesterId)
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}

func (r *Repository) FindCourseById(id uint) (*models.Course, error) {
	data := &models.Course{}
	err := r.db.Model(models.Course{}).Where("id = ?", id).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateCourse(id uint, req *models.CourseRequest) (*models.Course, error) {
	data := &models.Course{}
	err := r.db.Model(&models.Course{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        req.Name,
		"code":        req.Code,
		"credits":     req.Credits,
		"semester_id": req.SemesterId,
	}).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteCourse(id uint) error {
	return r.db.Model(&models.Course{}).Where("id = ?", id).Delete(&models.Course{}).Error
}
