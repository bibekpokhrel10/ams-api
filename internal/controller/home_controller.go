package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (server *Server) homePage(ctx *gin.Context) {
	logrus.Error("Welcome to Attendance Management System")
	ctx.JSON(http.StatusOK, "Welcome to Attendance Management System")
}
