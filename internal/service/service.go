package service

import "github.com/ams-api/internal/repository"

type IService interface {
	IUser
	IDepartment
	IClass
	ICourse
	ISemester
	IAttendance
	IEnrollment
}

type Service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return Service{
		repo: repo,
	}
}
