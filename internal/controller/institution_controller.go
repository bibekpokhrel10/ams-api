package controller

import (
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

func (server *Server) createInstitution(ctx *gin.Context) {
	var req *models.InstitutionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data, err := server.service.CreateInstitution(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) listInstitution(ctx *gin.Context) {
	req := &models.ListInstitutionRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	req.Prepare()
	datas, count, err := server.service.ListInstitution(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas, response.WithPagination(count, int(req.Page), int(req.Size))))
}

func (server *Server) getInstitutionById(ctx *gin.Context) {
	id := ctx.Param("id")
	moduleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetInstitutionById(uint(moduleId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateInstitution(ctx *gin.Context) {
	id := ctx.Param("id")
	moduleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.InstitutionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data, err := server.service.UpdateInstitution(uint(moduleId), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteInstitution(ctx *gin.Context) {
	id := ctx.Param("id")
	moduleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	err = server.service.DeleteInstitution(uint(moduleId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
