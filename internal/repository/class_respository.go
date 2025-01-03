package repository

import (
	"github.com/ams-api/internal/models"
)

type IClass interface {
	CreateClass(data *models.Class) (*models.Class, error)
	FindAllClass(req *models.ListClassRequest) (*[]models.Class, error)
	FindClassById(id uint) (*models.Class, error)
	UpdateClass(id uint, req *models.ClassRequest) (*models.Class, error)
	FindClassByInstructor(instructorId uint) (*[]models.Class, error)
	FindClassByInstitutionId(institutionId uint) (*[]models.Class, error)
	DeleteClass(id uint) error
}

func (r *Repository) CreateClass(data *models.Class) (*models.Class, error) {
	err := r.db.Model(&models.Class{}).Create(data).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindAllClass(req *models.ListClassRequest) (*[]models.Class, error) {
	datas := &[]models.Class{}
	f := r.db.Model(&models.Class{})
	if req.InstructorId != 0 {
		f = f.Where("instructor_id = ?", req.InstructorId)
	}
	err := f.Order("id desc").Preload("Course").Preload("User").Preload("Course.Semester").Preload("Course.Semester.Program").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindClassById(id uint) (*models.Class, error) {
	data := &models.Class{}
	err := r.db.Model(models.Class{}).Where("id = ?", id).Preload("Course").Preload("User").Preload("Course.Semester").Preload("Course.Semester.Program").Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (r *Repository) UpdateClass(id uint, req *models.ClassRequest) (*models.Class, error) {
	data := &models.Class{}
	err := r.db.Model(&models.Class{}).Where("id = ?", id).Updates(map[string]interface{}{
		"course_id":     req.CourseId,
		"year":          req.Year,
		"schedule":      req.Schedule,
		"instructor_id": req.InstructorID,
	}).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) FindClassByInstructor(instructorId uint) (*[]models.Class, error) {
	datas := &[]models.Class{}
	err := r.db.Model(&models.Class{}).Where("instructor_id = ?", instructorId).Preload("Course").Preload("User").Preload("Course.Semester").Preload("Course.Semester.Program").Find(datas).Error
	if err != nil {
		return nil, err
	}
	return datas, err
}

func (r *Repository) FindClassByInstitutionId(institutionId uint) (*[]models.Class, error) {
	datas := &[]models.Class{}
	err := r.db.Model(&models.Class{}).
		Joins("JOIN courses ON classes.course_id = courses.id").
		Joins("JOIN semesters ON courses.semester_id = semesters.id").
		Joins("JOIN programs ON semesters.program_id = programs.id").
		Where("programs.institution_id = ?", institutionId).
		Preload("Course").
		Preload("User").
		Preload("Course.Semester").
		Preload("Course.Semester.Program").
		Find(datas).Error

	if err != nil {
		return nil, err
	}
	return datas, nil
}

func (r *Repository) DeleteClass(id uint) error {
	return r.db.Model(&models.Class{}).Where("id = ?", id).Delete(&models.Class{}).Error
}
