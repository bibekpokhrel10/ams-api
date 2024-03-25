package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	gorm.Model
	Date      time.Time `json:"date"`
	ClassID   uint      `json:"class_id"`
	StudentID uint      `json:"student_id"`
	IsPresent bool      `json:"is_present"` // e.g., present, absent, late
	Remarks   string    `json:"remarks"`
}

type AttendanceResponse struct {
	Id        uint      `json:"id"`
	Date      time.Time `json:"date"`
	ClassID   uint      `json:"class_id"`
	StudentID uint      `json:"student_id"`
	IsPresent bool      `json:"is_present"`
	Remarks   string    `json:"remarks"`
}

type AttendanceRequest struct {
	Date      time.Time `json:"date"`
	ClassID   uint      `json:"class_id"`
	StudentID uint      `json:"student_id"`
	IsPresent bool      `json:"is_present"`
	Remarks   string    `json:"remarks"`
}

func (a *Attendance) AttendanceResponse() *AttendanceResponse {
	return &AttendanceResponse{
		Id:        a.ID,
		Date:      a.Date,
		ClassID:   a.ClassID,
		StudentID: a.StudentID,
		IsPresent: a.IsPresent,
		Remarks:   a.Remarks,
	}
}

func NewAttendance(req *AttendanceRequest) (*Attendance, error) {
	attendance := &Attendance{
		Date:      req.Date,
		ClassID:   req.ClassID,
		StudentID: req.StudentID,
		IsPresent: req.IsPresent,
		Remarks:   req.Remarks,
	}
	return attendance, nil
}
