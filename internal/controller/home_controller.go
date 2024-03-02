package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Welcome to Attendance Management System")
}
