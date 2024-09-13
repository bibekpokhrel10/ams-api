package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ams-api/internal/models"
	"github.com/ams-api/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req *models.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logrus.Error("error while creating user :: ", err)
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	isAvailable := server.service.CheckEmailAvailable(req.Email)
	if isAvailable {
		ctx.JSON(http.StatusConflict, response.ERROR(errors.New("email already exists")))
		return
	}
	user, err := server.service.CreateUser(req)
	if err != nil {
		logrus.Error("error while registering user :: ", err)
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(user))
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req *models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	res, code, err := server.service.LoginUser(req)
	if err != nil {
		ctx.JSON(code, response.ERROR(err))
		return
	}
	accessToken, err := server.tokenMaker.CreateToken(res.User.Id)
	res.Token = accessToken
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(res))
}

func (server *Server) getUserFromToken(ctx *gin.Context) {
	payload, err := server.getAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ERROR(err))
		return
	}
	data, err := server.service.GetUserById(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data.Role = "admin"
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) getAllUsers(ctx *gin.Context) {
	datas, err := server.service.ListAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas))
}

func (server *Server) activateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	moduleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.ActivateDeactivateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	res, err := server.service.ActivateDeactivateUser(uint(moduleId), req.IsActive)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(res))
}
