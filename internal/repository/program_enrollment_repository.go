package repository

import (
	"github.com/ams-api/internal/models"
)

type IProgramEnrollment interface {
	CreateProgramEnrollment(data *models.ProgramEnrollment) (*models.ProgramEnrollment, error)
	FindAllProgramEnrollment(req *models.ListProgramEnrollmentRequest) ([]models.ProgramEnrollment, int, error)
	FindProgramEnrollmentById(id uint) (*models.ProgramEnrollment, error)
	UpdateProgramEnrollment(id uint, req *models.ProgramEnrollmentRequest) (*models.ProgramEnrollment, error)
	FindProgramEnrollmentByProgramIdAndStudentId(programId, studentId uint) (*models.ProgramEnrollment, error)
	FindProgramEnrollmentByProgramId(programId uint) ([]models.ProgramEnrollment, error)
	DeleteProgramEnrollment(id uint) error
}

func (r *Repository) CreateProgramEnrollment(data *models.ProgramEnrollment) (*models.ProgramEnrollment, error) {
	err := r.db.Model(&models.ProgramEnrollment{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllProgramEnrollment(req *models.ListProgramEnrollmentRequest) ([]models.ProgramEnrollment, int, error) {
	var datas []models.ProgramEnrollment
	var count int64

	// Base query
	query := r.db.Model(&models.ProgramEnrollment{})

	// Apply filters
	if req.ProgramId != 0 {
		query = query.Where("program_enrollments.program_id = ?", req.ProgramId)
	}
	if req.StudentId != 0 {
		query = query.Where("program_enrollments.student_id = ?", req.StudentId)
	}

	// Count total records (before pagination)
	countQuery := query.Joins("JOIN users ON users.id = program_enrollments.student_id").
		Where("LOWER(users.email) LIKE LOWER(?)", "%"+req.Query+"%")
	err := countQuery.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	// Get paginated records with preloaded User
	err = query.Preload("User").
		Joins("JOIN users ON users.id = program_enrollments.student_id").
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

func (r *Repository) FindProgramEnrollmentById(id uint) (*models.ProgramEnrollment, error) {
	data := &models.ProgramEnrollment{}
	err := r.db.Model(models.ProgramEnrollment{}).Where("id = ?", id).Preload("User").Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) FindProgramEnrollmentByProgramId(programId uint) ([]models.ProgramEnrollment, error) {
	data := []models.ProgramEnrollment{}
	err := r.db.Model(models.ProgramEnrollment{}).Where("program_id = ?", programId).Preload("User").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateProgramEnrollment(id uint, req *models.ProgramEnrollmentRequest) (*models.ProgramEnrollment, error) {
	data := &models.ProgramEnrollment{}
	err := r.db.Model(&models.ProgramEnrollment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"student_id": req.StudentId,
		"program_id": req.ProgramId,
	}).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteProgramEnrollment(id uint) error {
	return r.db.Model(&models.ProgramEnrollment{}).Where("id = ?", id).Delete(&models.ProgramEnrollment{}).Error
}

func (r *Repository) FindProgramEnrollmentByProgramIdAndStudentId(programId, studentId uint) (*models.ProgramEnrollment, error) {
	data := &models.ProgramEnrollment{}
	err := r.db.Model(&models.ProgramEnrollment{}).Where("program_id = ?", programId).Where("student_id = ?", studentId).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}
