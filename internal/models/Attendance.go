package models

import (
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
	Status    string    `json:"status"`     // e.g., "present", "absent", "late"
	Remarks   string    `json:"remarks"`
}

type AttendanceRecordRequest struct {
	ClassId    uint                `json:"class_id"`
	Date       time.Time           `json:"date"`
	Attendance []StudentAttendance `json:"attendance"`
}

// StudentAttendance represents the attendance status for a single student
type StudentAttendance struct {
	StudentId uint   `json:"student_id"`
	Status    string `json:"status"` // e.g., "present", "absent", "late"
}

type AttendanceResponse struct {
	Id        uint           `json:"id"`
	Date      time.Time      `json:"date"`
	ClassID   uint           `json:"class_id"`
	Class     *ClassResponse `json:"class"`
	StudentID uint           `json:"student_id"`
	Student   *UserResponse  `json:"user_response"`
	IsPresent bool           `json:"is_present"`
	Status    string         `json:"status"`
	Remarks   string         `json:"remarks"`
}

type AttendanceRequest struct {
	Date      time.Time `json:"date"`
	ClassId   uint      `json:"class_id"`
	StudentId uint      `json:"student_id"`
	IsPresent bool      `json:"is_present"`
	Status    string    `json:"status"`
	Remarks   string    `json:"remarks"`
}

type ListAttendanceRequest struct {
	ListRequest
	Date      string `form:"date"`
	ClassId   uint   `form:"class_id"`
	StudentId uint   `form:"student_id"`
}

type AttendanceStats struct {
	TotalPresent  int `json:"total_present"`
	TotalAbsent   int `json:"total_absent"`
	TotalLate     int `json:"total_late"`
	TotalStudents int `json:"total_students"`
	TotalOnLeave  int `json:"total_on_leave"`
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
		Status:    a.Status,
		Remarks:   a.Remarks,
	}
}

func NewAttendance(req *AttendanceRequest) (*Attendance, error) {
	attendance := &Attendance{
		Date:      req.Date,
		ClassId:   req.ClassId,
		StudentId: req.StudentId,
		IsPresent: req.IsPresent,
		Status:    req.Status,
		Remarks:   req.Remarks,
	}
	return attendance, nil
}

func (req *AttendanceRecordRequest) Validate() error {
	return nil
}

func (req *AttendanceRecordRequest) Prepare() {

}

type StudentAttendanceAlert struct {
	StudentID                 int      `json:"student_id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	Email                     string   `json:"email"`
	CourseName                string   `json:"course_name"`
	CourseCode                string   `json:"course_code"`
	InstructorName            string   `json:"instructor_name"`
	ClassesAttended           int      `json:"classes_attended"`
	TotalClasses              int      `json:"total_classes"`
	AttendancePercentage      float64  `json:"attendance_percentage"`
	RecentAttendance          []string `json:"recent_attendance"`
	ClassesNeededForThreshold int      `json:"classes_needed_for_threshold"`
}

type SendAlertRequest struct {
	Threshold float64 `json:"threshold"`
	UserType  string  `json:"user_type"`
}
