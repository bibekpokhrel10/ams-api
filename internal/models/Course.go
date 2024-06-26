package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Credits    int       `json:"credits"`
	SemesterId uint      `json:"semester_id"`
	Semester   *Semester `gorm:"foreignKey:SemesterId" json:"semester"`
}

type CourseRequest struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Credits    int    `json:"credits"`
	SemesterId uint   `json:"semester_id"`
}

type CourseResponse struct {
	Id         uint              `json:"id"`
	Code       string            `json:"code"`
	Name       string            `json:"name"`
	Credits    int               `json:"credits"`
	SemesterId uint              `json:"semester_id"`
	Semester   *SemesterResponse `json:"semester"`
}

func (c *Course) CourseResponse() *CourseResponse {
	semesterResp := &SemesterResponse{}
	if c.Semester != nil {
		semesterResp = c.Semester.SemesterResponse()
	}
	return &CourseResponse{
		Id:         c.ID,
		Code:       c.Code,
		Name:       c.Name,
		Credits:    c.Credits,
		SemesterId: c.SemesterId,
		Semester:   semesterResp,
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

func (req *CourseRequest) Validate() error {
	if req.Name == "" {
		return errors.New("program name is required")
	}
	return nil
}

func (req *CourseRequest) Prepare() {
	req.Name = strings.TrimSpace(req.Name)
}
