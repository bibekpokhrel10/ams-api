package main

import (
	_ "time/tzdata"

	"github.com/ams-api/internal/config"
	"github.com/ams-api/internal/config/database"
	"github.com/ams-api/internal/controller"
	"github.com/ams-api/internal/repository"
	"github.com/ams-api/internal/service"
	"github.com/sirupsen/logrus"
)

// @IGaming-Main-API
// @version 1.0
// @description This is a iGaming-Main-API server.
// @termsOfService http://swagger.io/terms/

// @contact.email info.igaming.com
// @securityDefinitions.apikey ApiKeyAuth
// @in Header
// @name Authorization

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal("Error on loading configuration :: ", err)
	}
	db, err := database.NewPostgres(config)
	if err != nil {
		logrus.Fatal("Error on loading configuration :: ", err)
	}
	err = repository.AutoMigrate(db)
	if err != nil {
		logrus.Fatal("Error on migrating table :: ", err)
	}
	// err = seeder.Load(db)
	// if err != nil {
	// 	logrus.Fatal("Unable to load seeder :: ", err)
	// }
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	server, err := controller.NewServer(config, svc)
	if err != nil {
		logrus.Fatal("Cannot create server : ", err)
	}
	err = server.Start(config.PORT)
	if err != nil {
		logrus.Fatal("Unable to start server ", err)
	}
}
