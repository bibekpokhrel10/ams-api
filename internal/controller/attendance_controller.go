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

func (server *Server) createAttendance(ctx *gin.Context) {
	var req *models.AttendanceRecordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	err := server.service.CreateAttendance(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success("successfully created attendance record"))
}

func (server *Server) listAttendance(ctx *gin.Context) {
	var req *models.ListAttendanceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	req.Prepare()
	datas, count, err := server.service.ListAttendance(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(datas, response.WithPagination(count, int(req.Page), int(req.Size))))
}

func (server *Server) listAttendanceStats(ctx *gin.Context) {
	date := ctx.Query("date")
	if date == "" {
		logrus.Error("error while getting attendance stats :: date is required")
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("date is required")))
		return
	}
	classIdReq := ctx.Query("class_id")
	if classIdReq == "" {
		logrus.Error("error while getting attendance stats :: class_id is required")
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("class_id is required")))
		return
	}
	classId, err := strconv.Atoi(classIdReq)
	if err != nil {
		logrus.Error("error while getting class id :: ", err)
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data, err := server.service.GetAttendanceStatsByDate(uint(classId), date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) getAttendanceById(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	data, err := server.service.GetAttendanceById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) updateAttendance(ctx *gin.Context) {
	id := ctx.Param("id")
	moduleId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	var req *models.AttendanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	data, err := server.service.UpdateAttendance(uint(moduleId), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(data))
}

func (server *Server) deleteAttendance(ctx *gin.Context) {
	id := ctx.Param("id")
	depId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ERROR(err))
		return
	}
	Attendance, err := server.service.GetAttendanceById(uint(depId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(errors.New("attendance not found")))
		return
	}
	err = server.service.DeleteAttendance(Attendance.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(""))
}

func (server *Server) SendAttendanceAlert(ctx *gin.Context) {
	server.service.SendAttendanceAlert()
}

func (server *Server) SendAttendanceAlertManually(ctx *gin.Context) {
	var req *models.SendAlertRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ERROR(err))
		return
	}
	payload, err := server.getAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ERROR(err))
		return
	}
	server.service.SendAttendanceAlertAccordingToUserType(payload.UserId, req)
	ctx.JSON(http.StatusOK, response.Success(""))
}
