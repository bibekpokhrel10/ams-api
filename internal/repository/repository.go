package repository

import (
	"github.com/ams-api/internal/models"
	"github.com/jinzhu/gorm"
)

type IRepository interface {
	IUser
	IDepartment
	IAttendance
	IClass
	ICourse
	ISemester
	IEnrollment
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
		models.Department{},
	).Error
}
