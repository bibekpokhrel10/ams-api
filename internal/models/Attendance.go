package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Attendance struct {
	gorm.Model
	Date      time.Time `json:"date"`
	ClassId   uint      `json:"class_id"`
	Class     *Class    `gorm:"foreignKey:ClassId" json:"class"`
	StudentId uint      `json:"student_id"`
	Student   *User     `gorm:"foreignKey:StudentId" json:"student"`
	IsPresent bool      `json:"is_present"` // e.g., present, absent, late
	Remarks   string    `json:"remarks"`
}

type AttendanceResponse struct {
	Id        uint           `json:"id"`
	Date      time.Time      `json:"date"`
	ClassID   uint           `json:"class_id"`
	Class     *ClassResponse `json:"class"`
	StudentID uint           `json:"student_id"`
	Student   *UserResponse  `json:"user_response"`
	IsPresent bool           `json:"is_present"`
	Remarks   string         `json:"remarks"`
}

type AttendanceRequest struct {
	Date      time.Time `json:"date"`
	ClassId   uint      `json:"class_id"`
	StudentId uint      `json:"student_id"`
	IsPresent bool      `json:"is_present"`
	Remarks   string    `json:"remarks"`
}

func (a *Attendance) AttendanceResponse() *AttendanceResponse {
	studentResp := &UserResponse{}
	if a.Student != nil {
		studentResp = a.Student.UserResponse()
	}
	classResp := &ClassResponse{}
	if a.Class != nil {
		classResp = a.Class.ClassResponse()
	}
	return &AttendanceResponse{
		Id:        a.ID,
		Date:      a.Date,
		ClassID:   a.ClassId,
		Class:     classResp,
		StudentID: a.StudentId,
		Student:   studentResp,
		IsPresent: a.IsPresent,
		Remarks:   a.Remarks,
	}
}

func NewAttendance(req *AttendanceRequest) (*Attendance, error) {
	attendance := &Attendance{
		Date:      req.Date,
		ClassId:   req.ClassId,
		StudentId: req.StudentId,
		IsPresent: req.IsPresent,
		Remarks:   req.Remarks,
	}
	return attendance, nil
}

func (req *AttendanceRequest) Validate() error {
	return nil
}

func (req *AttendanceRequest) Prepare() {
	req.Remarks = strings.TrimSpace(req.Remarks)
}
