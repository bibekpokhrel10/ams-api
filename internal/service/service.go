package service

import "github.com/ams-api/internal/repository"

type IService interface {
	IUser
	IClass
	ICourse
	ISemester
	IAttendance
	IEnrollment
	IProgram
}

type Service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return Service{
		repo: repo,
	}
}
