package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) createEnrollment(ctx *gin.Context) {
	var req *models.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	data, err := server.service.CreateEnrollment(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) listEnrollment(ctx *gin.Context) {
	datas, err := server.service.ListEnrollment()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas))
}

func (server *Server) getEnrollmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetEnrollmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateEnrollment(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.EnrollmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	Enrollment, err := server.service.GetEnrollmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("enrollment not found")))
		return
	}
	data, err := server.service.UpdateEnrollment(Enrollment.Id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteEnrollment(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Enrollment, err := server.service.GetEnrollmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("enrollment not found")))
		return
	}
	err = server.service.DeleteEnrollment(Enrollment.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
