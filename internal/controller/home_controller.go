package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Homepage godoc
// @Summary      Show Computeshere home page message
// @Tags         Home
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [get]
func (server *Server) homePage(ctx *gin.Context) {
	logrus.Error("Welcome to Attendance Management System")
	ctx.JSON(http.StatusOK, "Welcome to Attendance Management System")
}
