package controller

import (
	"net/http"

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
