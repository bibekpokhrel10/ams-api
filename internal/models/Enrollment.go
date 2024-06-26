package models

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentId uint `json:"student_id"`
	ClassId   uint `json:"class_id"`
	Year      int  `json:"year"`
	Duration  int  `json:"duration"`
}

type EnrollmentResponse struct {
	Id        uint `json:"id"`
	StudentId uint `json:"student_id"`
	ClassId   uint `json:"class_id"`
}

type EnrollmentRequest struct {
	StudentId uint `json:"student_id"`
	ClassId   uint `json:"class_id"`
}

func NewEnrollment(req *EnrollmentRequest) (*Enrollment, error) {
	enrollment := &Enrollment{
		StudentId: req.StudentId,
		ClassId:   req.ClassId,
	}
	return enrollment, nil
}

func (e *Enrollment) EnrollmentResponse() *EnrollmentResponse {
	return &EnrollmentResponse{
		Id:        e.ID,
		StudentId: e.StudentId,
		ClassId:   e.ClassId,
	}
}
