package models

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Code       string `json:"code"`
	Name       string `json:"name"`
	Credits    int    `json:"credits"`
	SemesterId int    `json:"semester_id"`
}

type CourseRequest struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Credits    int    `json:"credits"`
	SemesterId int    `json:"semester_id"`
}

type CourseResponse struct {
	Id         uint   `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Credits    int    `json:"credits"`
	SemesterId int    `json:"semester_id"`
}

func (c *Course) CourseResponse() *CourseResponse {
	return &CourseResponse{
		Id:         c.ID,
		Code:       c.Code,
		Name:       c.Name,
		Credits:    c.Credits,
		SemesterId: c.SemesterId,
	}
}

func NewCourse(req *CourseRequest) (*Course, error) {
	course := &Course{
		Code:       req.Code,
		Name:       req.Name,
		Credits:    req.Credits,
		SemesterId: req.SemesterId,
	}
	return course, nil
}
