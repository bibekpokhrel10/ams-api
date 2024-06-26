package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	CourseId     uint    `json:"course_id"`
	Course       *Course `gorm:"foreignKey:CourseId" json:"course"`
	Year         int     `json:"year"`
	Schedule     string  `json:"schedule"`
	InstructorID uint    `json:"instructor_id"`
	Instructor   *User   `gorm:"foreignKey:InstructorID" json:"instructor"`
}

type ClassResponse struct {
	Id           uint            `json:"id"`
	CourseId     uint            `json:"course_id"`
	Course       *CourseResponse `json:"course"`
	Year         int             `json:"year"`
	Schedule     string          `json:"schedule"`
	InstructorID uint            `json:"instructor_id"`
	Instructor   *UserResponse   `json:"instructor"`
}

type ClassRequest struct {
	CourseId     uint   `json:"course_id"`
	Year         int    `json:"year"`
	Schedule     string `json:"schedule"`
	InstructorID uint   `json:"instructor_id"`
}

func (c *Class) ClassResponse() *ClassResponse {
	courseResp := &CourseResponse{}
	if c.Course != nil {
		courseResp = c.Course.CourseResponse()
	}
	instructorResp := &UserResponse{}
	if c.Instructor != nil {
		instructorResp = c.Instructor.UserResponse()
	}
	return &ClassResponse{
		Id:           c.ID,
		CourseId:     c.CourseId,
		Course:       courseResp,
		Year:         c.Year,
		Schedule:     c.Schedule,
		InstructorID: c.InstructorID,
		Instructor:   instructorResp,
	}
}

func NewClass(req *ClassRequest) (*Class, error) {
	class := &Class{
		CourseId:     req.CourseId,
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
