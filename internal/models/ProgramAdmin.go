package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProgramAdmin struct {
	gorm.Model
	ProgramId uint     `json:"program_id"`
	Program   *Program `gorm:"foreignKey:ProgramId" json:"program"`
	UserId    uint     `json:"user_id"`
	User      *User    `gorm:"foreignKey:UserId" json:"user"`
}

type ProgramAdminResponse struct {
	Id        uint             `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	ProgramId uint             `json:"program_id"`
	Program   *ProgramResponse `json:"program"`
	UserId    uint             `json:"user_id"`
	User      *UserResponse    `json:"user"`
}

type ProgramAdminRequest struct {
	ProgramId uint `json:"program_id"`
	UserId    uint `json:"user_id"`
}

type ListProgramAdminRequest struct {
	ListRequest
	ProgramId uint `form:"program_id"`
}

func (p *ProgramAdmin) ProgramAdminResponse() *ProgramAdminResponse {
	return &ProgramAdminResponse{
		Id:        p.ID,
		CreatedAt: p.CreatedAt,
		ProgramId: p.ProgramId,
		Program:   p.Program.ProgramResponse(),
		UserId:    p.UserId,
		User:      p.User.UserResponse(),
	}
}

func NewProgramAdmin(req *ProgramAdminRequest) (*ProgramAdmin, error) {
	programAdmin := &ProgramAdmin{
		ProgramId: req.ProgramId,
		UserId:    req.UserId,
	}
	return programAdmin, nil
}

func (p *ProgramAdminRequest) Validate() error {
	// if p.Name == "" {
	// 	return errors.New("program name is required")
	// }
	// if p.Type == "" {
	// 	return errors.New("program type is required")
	// } else {
	// 	if p.Type != "undergraduate" && p.Type != "graduate" {
	// 		return errors.New("invalid program type, must be undergraduate or graduate")
	// 	}
	// }
	return nil
}

func (p *ProgramAdminRequest) Prepare() {

}
