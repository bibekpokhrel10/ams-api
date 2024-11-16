package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model
	InstitutionId uint   `json:"institution_id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Duration      string `json:"duration"`
}

type ProgramResponse struct {
	Id            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	InstitutionId uint      `json:"institution_id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
}

type ProgramRequest struct {
	InstitutionId uint   `json:"institution_id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
}

type ListProgramRequest struct {
	ListRequest
	InstitutionId uint `form:"institution_id"`
}

func (p *Program) ProgramResponse() *ProgramResponse {
	return &ProgramResponse{
		Id:            p.ID,
		CreatedAt:     p.CreatedAt,
		InstitutionId: p.InstitutionId,
		Name:          p.Name,
		Type:          p.Type,
	}
}

func NewProgram(req *ProgramRequest) (*Program, error) {
	program := &Program{
		InstitutionId: req.InstitutionId,
		Name:          req.Name,
		Type:          req.Type,
	}
	return program, nil
}

func (p *ProgramRequest) Validate() error {
	if p.Name == "" {
		return errors.New("program name is required")
	}
	// if p.Type == "" {
	// 	return errors.New("program type is required")
	// } else {
	// 	if p.Type != "undergraduate" && p.Type != "graduate" {
	// 		return errors.New("invalid program type, must be undergraduate or graduate")
	// 	}
	// }
	return nil
}

func (p *ProgramRequest) Prepare() {
	p.Name = strings.TrimSpace(p.Name)
	p.Type = strings.TrimSpace(strings.ToLower(p.Type))
}
