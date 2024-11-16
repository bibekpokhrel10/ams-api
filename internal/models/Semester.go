package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Semester struct {
	gorm.Model
	Name       string   `json:"name"`
	Year       string   `json:"year"`
	ProgramId  uint     `json:"program_id"`
	Program    *Program `gorm:"foreignKey:ProgramId" json:"program"`
	TimePeriod string   `json:"time_period"`
}

type SemesterResponse struct {
	Id         uint             `json:"id"`
	CreatedAt  time.Time        `json:"created_at"`
	Name       string           `json:"name"`
	Year       string           `json:"year"`
	ProgramId  uint             `json:"program_id"`
	Program    *ProgramResponse `json:"program"`
	TimePeriod string           `json:"time_period"`
}

type SemesterRequest struct {
	Name       string `json:"name"`
	Year       string `json:"year"`
	ProgramId  uint   `json:"program_id"`
	TimePeriod string `json:"time_period"`
}

type ListSemesterRequest struct {
	ListRequest
	ProgramId uint `form:"program_id"`
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
	programResp := &ProgramResponse{}
	if s.Program != nil {
		programResp = s.Program.ProgramResponse()
	}
	return &SemesterResponse{
		Id:         s.ID,
		CreatedAt:  s.CreatedAt,
		Name:       s.Name,
		ProgramId:  s.ProgramId,
		Year:       s.Year,
		Program:    programResp,
		TimePeriod: s.TimePeriod,
	}
}

func (req *SemesterRequest) Validate() error {
	if req.Name == "" {
		return errors.New("program name is required")
	}
	return nil
}

func (req *SemesterRequest) Prepare() {
	req.Name = strings.TrimSpace(req.Name)
	req.Year = strings.TrimSpace(req.Year)
	req.TimePeriod = strings.TrimSpace(req.TimePeriod)
}
