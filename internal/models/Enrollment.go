package models

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentID uint `json:"student_id"`
	ClassID   uint `json:"class_id"`
}

type EnrollmentResponse struct {
	Id        uint `json:"id"`
	StudentID uint `json:"student_id"`
	ClassID   uint `json:"class_id"`
}

type EnrollmentRequest struct {
	StudentID uint `json:"student_id"`
	ClassID   uint `json:"class_id"`
}

func NewEnrollment(req *EnrollmentRequest) (*Enrollment, error) {
	enrollment := &Enrollment{
		StudentID: req.StudentID,
		ClassID:   req.ClassID,
	}
	return enrollment, nil
}

func (e *Enrollment) EnrollmentResponse() *EnrollmentResponse {
	return &EnrollmentResponse{
		Id:        e.ID,
		StudentID: e.StudentID,
		ClassID:   e.ClassID,
	}
}
