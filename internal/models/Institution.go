package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Institution struct {
	gorm.Model
	Name string `json:"name"`
}

type InstitutionResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type InstitutionRequest struct {
	Name string `json:"name"`
}

type ListInstitutionRequest struct {
	ListRequest
	Name string `form:"name"`
}

func (p *Institution) InstitutionResponse() *InstitutionResponse {
	return &InstitutionResponse{
		Id:        p.ID,
		CreatedAt: p.CreatedAt,
		Name:      p.Name,
	}
}

func NewInstitution(req *InstitutionRequest) *Institution {
	Institution := &Institution{
		Name: req.Name,
	}
	return Institution
}

func (p *InstitutionRequest) Validate() error {
	if p.Name == "" {
		return errors.New("institution name is required")
	}
	return nil
}

func (p *InstitutionRequest) Prepare() {
	p.Name = strings.TrimSpace(p.Name)
}
