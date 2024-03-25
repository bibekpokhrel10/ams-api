package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) createCourse(ctx *gin.Context) {
	var req *models.CourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	data, err := server.service.CreateCourse(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) listCourse(ctx *gin.Context) {
	datas, err := server.service.ListCourse()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas))
}

func (server *Server) getCourseById(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetCourseById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.CourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	Course, err := server.service.GetCourseById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("course not found")))
		return
	}
	data, err := server.service.UpdateCourse(Course.Id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Course, err := server.service.GetCourseById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("course not found")))
		return
	}
	err = server.service.DeleteCourse(Course.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
