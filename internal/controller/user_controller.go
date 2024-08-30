package controller

import (
	"errors"
	"net/http"

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
