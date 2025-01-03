package main

import (
	"fmt"
	"path"
	"runtime"
	_ "time/tzdata"

	"github.com/ams-api/internal/config"
	"github.com/ams-api/internal/config/database"
	"github.com/ams-api/internal/controller"
	"github.com/ams-api/internal/repository"
	"github.com/ams-api/internal/seeder"
	"github.com/ams-api/internal/service"
	"github.com/sirupsen/logrus"
)

// @IGaming-Main-API
// @version 1.0
// @description This is a iGaming-Main-API server.
// @termsOfService http://swagger.io/terms/

// @contact.email info.ams.com
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
	if err := seeder.LoadSeeder(db); err != nil {
		logrus.Fatal("Unable to load seeder :: ", err)
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
	// 		fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	// 		//return frame.Function, fileName
	// 		return "", fileName
	// 	},
	// })
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableQuote:    true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("\x1b[36m%s:%d\x1b[0m", filename, f.Line)
		},
	})
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
