package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) createClass(ctx *gin.Context) {
	var req *models.ClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	data, err := server.service.CreateClass(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) listClass(ctx *gin.Context) {
	datas, err := server.service.ListClass()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas))
}

func (server *Server) getClassById(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetClassById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateClass(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.ClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	// if err := req.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ERROR(err))
	// 	return
	// }
	Class, err := server.service.GetClassById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("class not found")))
		return
	}
	data, err := server.service.UpdateClass(Class.Id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteClass(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Class, err := server.service.GetClassById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("class not found")))
		return
	}
	err = server.service.DeleteClass(Class.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
