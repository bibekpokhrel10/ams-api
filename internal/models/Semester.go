package models

import (
	"github.com/jinzhu/gorm"
)

type Semester struct {
	gorm.Model
	Name       string `json:"name"`
	Year       string `json:"year"`
	ProgramId  uint   `json:"program_id"`
	TimePeriod string `json:"time_period"`
}

type SemesterResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Year       string `json:"year"`
	ProgramId  uint   `json:"program_id"`
	TimePeriod string `json:"time_period"`
}

type SemesterRequest struct {
	Name       string `json:"name"`
	Year       string `json:"year"`
	ProgramId  uint   `json:"program_id"`
	TimePeriod string `json:"time_period"`
}

func NewSemester(req *SemesterRequest) (*Semester, error) {
	semester := &Semester{
		Name:       req.Name,
		Year:       req.Year,
		ProgramId:  req.ProgramId,
		TimePeriod: req.TimePeriod,
	}
	return semester, nil
}

func (s *Semester) SemesterResponse() *SemesterResponse {
	return &SemesterResponse{
		Id:         s.ID,
		Name:       s.Name,
		ProgramId:  s.ProgramId,
		Year:       s.Year,
		TimePeriod: s.TimePeriod,
	}
}
