package models

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model
	StudentId uint  `json:"student_id"`
	ClassId   uint  `json:"class_id"`
	User      *User `gorm:"foreignKey:StudentId" json:"user"`
}

type EnrollmentResponse struct {
	Id        uint          `json:"id"`
	StudentId uint          `json:"student_id"`
	ClassId   uint          `json:"class_id"`
	User      *UserResponse `json:"user"`
}

type EnrollmentRequest struct {
	StudentId uint `json:"student_id"`
	ClassId   uint `json:"class_id"`
}

type EnrollmentRequestPayload struct {
	StudentIds []uint `json:"student_ids"`
	ClassId    uint   `json:"class_id"`
}

type ListEnrollmentRequest struct {
	ListRequest
	ClassId   uint `form:"class_id"`
	StudentId uint `form:"student_id"`
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
		User:      e.User.UserResponse(),
	}
}
