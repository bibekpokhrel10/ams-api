package database

import (
	"log"

	"github.com/ams-api/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

func NewPostgres(config config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to %s database: %v", config.DBDriver, err)
	}
	// sqlDB := db.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
