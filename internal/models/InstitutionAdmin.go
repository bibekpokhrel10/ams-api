package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type InstitutionAdmin struct {
	gorm.Model
	InstitutionId uint         `json:"institution_id"`
	Institution   *Institution `json:"institution" gorm:"foreignKey:InstitutionId"`
	UserId        uint         `json:"user_id"`
	User          *User        `json:"user" gorm:"foreignKey:UserId"`
}

type InstitutionAdminResponse struct {
	Id            uint                 `json:"id"`
	CreatedAt     time.Time            `json:"created_at"`
	UserId        uint                 `json:"user_id"`
	User          *UserResponse        `json:"user"`
	InstitutionId uint                 `json:"institution_id"`
	Institution   *InstitutionResponse `json:"institution"`
}

type InstitutionAdminRequest struct {
	InstitutionId uint `json:"institution_id"`
	UserId        uint `json:"user_id"`
}

type ListInstitutionAdminRequest struct {
	ListRequest
	InstitutionId uint `json:"institution_id"`
}

func (p *InstitutionAdmin) InstitutionAdminResponse() *InstitutionAdminResponse {
	return &InstitutionAdminResponse{
		Id:            p.ID,
		CreatedAt:     p.CreatedAt,
		UserId:        p.UserId,
		User:          p.User.UserResponse(),
		InstitutionId: p.InstitutionId,
		Institution:   p.Institution.InstitutionResponse(),
	}
}

func NewInstitutionAdmin(req *InstitutionAdminRequest) *InstitutionAdmin {
	InstitutionAdmin := &InstitutionAdmin{
		InstitutionId: req.InstitutionId,
		UserId:        req.UserId,
	}
	return InstitutionAdmin
}

func (p *InstitutionAdminRequest) Validate() error {
	return nil
}

func (p *InstitutionAdminRequest) Prepare() {

}
