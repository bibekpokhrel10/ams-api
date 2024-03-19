package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) createDepartment(ctx *gin.Context) {
	var req *models.DepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data, err := server.service.CreateDepartment(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) listDepartment(ctx *gin.Context) {
	datas, err := server.service.ListDepartment()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas))
}

func (server *Server) getDepartmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetDepartmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateDepartment(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.DepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	department, err := server.service.GetDepartmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("department not found")))
		return
	}
	data, err := server.service.UpdateDepartment(department.Id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteDepartment(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	department, err := server.service.GetDepartmentById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("department not found")))
		return
	}
	err = server.service.DeleteDepartment(department.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
