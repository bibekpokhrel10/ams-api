package repository

import (
	"github.com/ams-api/internal/models"
)

type IEnrollment interface {
	CreateEnrollment(data *models.Enrollment) (*models.Enrollment, error)
	FindAllEnrollment(req *models.ListEnrollmentRequest) ([]models.Enrollment, int, error)
	FindEnrollmentById(id uint) (*models.Enrollment, error)
	UpdateEnrollment(id uint, req *models.EnrollmentRequest) (*models.Enrollment, error)
	DeleteEnrollment(id uint) error
	FindProgramEnrollmentByClassIdAndStudentId(classId, studentId uint) (*models.Enrollment, error)
}

func (r *Repository) CreateEnrollment(data *models.Enrollment) (*models.Enrollment, error) {
	err := r.db.Model(&models.Enrollment{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllEnrollment(req *models.ListEnrollmentRequest) ([]models.Enrollment, int, error) {
	var datas []models.Enrollment
	var count int64

	// Base query
	query := r.db.Model(&models.Enrollment{})

	// Apply filters
	if req.ClassId != 0 {
		query = query.Where("enrollments.class_id = ?", req.ClassId)
	}
	if req.StudentId != 0 {
		query = query.Where("enrollments.student_id = ?", req.StudentId)
	}

	// Count total records (before pagination)
	countQuery := query.Joins("JOIN users ON users.id = enrollments.student_id").
		Where("LOWER(users.email) LIKE LOWER(?)", "%"+req.Query+"%")
	err := countQuery.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated records with preloaded User
	err = query.Preload("User").
		Joins("JOIN users ON users.id = enrollments.student_id").
		Where("LOWER(users.email) LIKE LOWER(?)", "%"+req.Query+"%").
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(int(req.Size)).
		Offset(int(req.Size * (req.Page - 1))).
		Find(&datas).Error
	if err != nil {
		return nil, 0, err
	}

	return datas, int(count), nil
}

func (r *Repository) FindEnrollmentById(id uint) (*models.Enrollment, error) {
	data := &models.Enrollment{}
	err := r.db.Model(models.Enrollment{}).Where("id = ?", id).Preload("User").Take(data).Error
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

func (r *Repository) FindProgramEnrollmentByClassIdAndStudentId(classId, studentId uint) (*models.Enrollment, error) {
	data := &models.Enrollment{}
	err := r.db.Model(&models.Enrollment{}).Where("classId = ?", classId).Where("student_id = ?", studentId).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
