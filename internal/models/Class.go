package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	CourseId     uint   `json:"course_id"`
	SemesterId   uint   `json:"semester_id"`
	Year         int    `json:"year"`
	Schedule     string `json:"schedule"`
	InstructorID uint   `json:"instructor_id"`
}

type ClassResponse struct {
	Id           uint   `json:"id"`
	CourseId     uint   `json:"course_id"`
	SemesterId   uint   `json:"semester_id"`
	Year         int    `json:"year"`
	Schedule     string `json:"schedule"`
	InstructorID uint   `json:"instructor_id"`
}

type ClassRequest struct {
	CourseId     uint   `json:"course_id"`
	SemesterId   uint   `json:"semester_id"`
	Year         int    `json:"year"`
	Schedule     string `json:"schedule"`
	InstructorID uint   `json:"instructor_id"`
}

func (c *Class) ClassResponse() *ClassResponse {
	return &ClassResponse{
		Id:           c.ID,
		CourseId:     c.CourseId,
		SemesterId:   c.SemesterId,
		Year:         c.Year,
		Schedule:     c.Schedule,
		InstructorID: c.InstructorID,
	}
}

func NewClass(req *ClassRequest) (*Class, error) {
	class := &Class{
		CourseId:     req.CourseId,
		SemesterId:   req.SemesterId,
		Year:         req.Year,
		Schedule:     req.Schedule,
		InstructorID: req.InstructorID,
	}
	return class, nil
}

func (req *ClassRequest) Validate() error {
	if req.Schedule == "" {
		return errors.New("program name is required")
	}
	return nil
}

func (req *ClassRequest) Prepare() {
	req.Schedule = strings.TrimSpace(req.Schedule)
}
