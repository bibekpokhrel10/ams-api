package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
)

// CreateAgentGame godoc
// @Summary Create Agent Game
// @Description Create Agent Game
// @Tags Agent Game
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param body body models.DepartmentRequest true "Create Agent Game"
// @Success 200 {object} models.DepartmentResponse
// @Router /agent_games [post]
func (server *Server) createDepartment(ctx *gin.Context) {
	var req *models.DepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	Department, err := server.service.CreateDepartment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(Department))
}

// ListAgentGame godoc
// @Summary List Agent Game
// @Description List Agent Game
// @Tags Agent Game
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.DepartmentResponse
// @Router /agent_games [get]
func (server *Server) listDepartment(ctx *gin.Context) {
	Department, err := server.service.ListDepartment()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(Department))
}

// GetDepartmentById godoc
// @Summary Get Agent Game
// @Description Get Agent Game from Id
// @Tags Agent Game
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Agent Games id"
// @Success 200 {object} models.DepartmentResponse
// @Router /agent_games/{id} [get]
func (server *Server) getDepartmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	agid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Department, err := server.service.GetDepartmentById(uint(agid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(Department))
}

// UpdateAgentGame godoc
// @Summary Update Agent Game
// @Description Update Agent Games with the input payload
// @Tags Agent Game
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Agent Games id"
// @Param body body models.DepartmentRequest true "Update AgentGame"
// @Success 200 {object} models.DepartmentResponse
// @Router /agent_games/{id} [put]
func (server *Server) updateDepartment(ctx *gin.Context) {
	id := ctx.Param("id")
	agid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.DepartmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	Department, err := server.service.GetDepartmentById(uint(agid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(errors.New("agent games not found")))
		return
	}
	data, err := server.service.UpdateDepartment(Department.Id, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

// DeleteDepartment godoc
// @Summary Delete Agent Game
// @Description Delete Agent Game with the input payload
// @Tags Agent Game
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Agent Games id"
// @Success 200
// @Router /agent_games/{id} [delete]
func (server *Server) deleteDepartment(ctx *gin.Context) {
	id := ctx.Param("id")
	agid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Department, err := server.service.GetDepartmentById(uint(agid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(errors.New("agent games not found")))
		return
	}
	err = server.service.DeleteDepartment(Department.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}
