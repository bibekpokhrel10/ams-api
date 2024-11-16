package repository

import (
	"github.com/ams-api/internal/models"
	"github.com/jinzhu/gorm"
)

type IRepository interface {
	IUser
	IAttendance
	IClass
	ICourse
	ISemester
	IEnrollment
	IProgram
	IInstitution
	IProgramEnrollment
}

// Repostory type struct
type Repository struct {
	db *gorm.DB
}

// NewRepository create new repository
func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db,
	}
}

// Migrate up database table
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.Attendance{},
		models.Class{},
		models.Course{},
		models.Semester{},
		models.Enrollment{},
		models.Program{},
		models.Institution{},
		models.ProgramEnrollment{},
	).Error
}
