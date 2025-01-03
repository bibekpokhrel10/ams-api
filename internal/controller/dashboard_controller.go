package controller

import (
	"net/http"
	"strconv"

	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) getDashboard(ctx *gin.Context) {
	payload, err := server.getAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ERROR(err))
		return
	}
	data, err := server.service.GetDashboard(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) getStudentClassAttendance(ctx *gin.Context) {
	payload, err := server.getAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ERROR(err))
		return
	}
	id := ctx.Param("id")
	classId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetStudentClassAttendanceDetails(payload.UserId, uint(classId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) getTeacherClassAttendance(ctx *gin.Context) {
	id := ctx.Param("id")
	classId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetTeacherClassAttendanceDetails(uint(classId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}
