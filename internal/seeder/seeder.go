package seeder

import (
	"github.com/ams-api/internal/models"
	"github.com/ams-api/util/password"
	"github.com/jinzhu/gorm"
)

func LoadSeeder(db *gorm.DB) error {
	err := LoadUserSeeder(db)
	if err != nil {
		return err
	}
	return nil
}

func LoadUserSeeder(db *gorm.DB) error {
	user := &models.User{
		FirstName:   "Admin",
		LastName:    "User",
		Email:       "admin@admin.com",
		IsAdmin:     true,
		Password:    "Admin@1234",
		IsActive:    true,
		DateOfBirth: "2000-01-01",
		Gender:      "others",
		UserType:    "super_admin",
	}
	hashpwd, err := password.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpwd
	err = db.Model(&models.User{}).Where(models.User{Email: "admin@admin.com"}).FirstOrCreate(user).Take(user).Error
	if err != nil {
		return err
	}
	return nil
}
