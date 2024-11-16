package models

import "github.com/jinzhu/gorm"

type ProgramEnrollment struct {
	gorm.Model
	StudentId uint  `json:"student_id"`
	ProgramId uint  `json:"program_id"`
	User      *User `gorm:"foreignKey:StudentId" json:"user"`
}

type ProgramEnrollmentResponse struct {
	Id        uint          `json:"id"`
	StudentId uint          `json:"student_id"`
	ProgramId uint          `json:"program_id"`
	User      *UserResponse `json:"user"`
}

type ProgramEnrollmentRequest struct {
	StudentId uint `json:"student_id"`
	ProgramId uint `json:"program_id"`
}

type ProgramEnrollmentRequestPayload struct {
	StudentIds []uint `json:"student_ids"`
	ProgramId  uint   `json:"program_id"`
}

type ListProgramEnrollmentRequest struct {
	ListRequest
	ProgramId uint `form:"program_id"`
	StudentId uint `form:"student_id"`
}

func NewProgramEnrollment(req *ProgramEnrollmentRequest) (*ProgramEnrollment, error) {
	programEnrollment := &ProgramEnrollment{
		StudentId: req.StudentId,
		ProgramId: req.ProgramId,
	}
	return programEnrollment, nil
}

func (e *ProgramEnrollment) ProgramEnrollmentResponse() *ProgramEnrollmentResponse {
	return &ProgramEnrollmentResponse{
		Id:        e.ID,
		StudentId: e.StudentId,
		ProgramId: e.ProgramId,
		User:      e.User.UserResponse(),
	}
}
