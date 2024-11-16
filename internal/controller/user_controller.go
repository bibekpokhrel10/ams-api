package controller

import (
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
		logrus.Error("error while login user :: ", err)
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
	var req models.ListUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	req.Prepare()
	datas, count, err := server.service.ListAllUsers(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas, response.WithPagination(count, int(req.Page), int(req.Size))))
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

func (server *Server) updateUserPassword(ctx *gin.Context) {
	var req *models.UpdateUserPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	id := ctx.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.UpdateUserPassword(uint(userId), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	err = server.service.DeleteUser(uint(userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success("successfully deleted user"))
}

func (server *Server) updateUserType(ctx *gin.Context) {
	payload, err := server.getAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ERROR(err))
		return
	}
	var req *models.UpdateUserType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	id := ctx.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.UpdateUserType(payload.UserId, uint(userId), req.UserType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}
