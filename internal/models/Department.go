package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name string `json:"name"`
	Type string `json:"type"` // teacher/student
}

type DepartmentRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DepartmentResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (d *Department) DepartmentResponse() *DepartmentResponse {
	return &DepartmentResponse{
		Id:   d.ID,
		Name: d.Name,
		Type: d.Type,
	}
}

func NewDepartment(req *DepartmentRequest) *Department {
	return &Department{
		Name: req.Name,
		Type: req.Type,
	}
}

func (r *DepartmentRequest) Validate() error {
	if r.Name == "" {
		return ErrDepartmentRequired
	}
	if r.Type == "" {
		return ErrDepartmentTypeRequired
	} else {
		if r.Type != "teacher" && r.Type != "student" {
			return ErrInvalidDepartmentType
		}
	}
	return nil
}
