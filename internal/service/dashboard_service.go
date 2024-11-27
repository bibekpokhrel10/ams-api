package service

import (
	"github.com/ams-api/internal/models"
	"github.com/ams-api/util/constants"
	"github.com/sirupsen/logrus"
)

type IDashboard interface {
	GetDashboard(userId uint) (*models.DashboardResponse, error)
}

func (s Service) GetDashboard(userId uint) (*models.DashboardResponse, error) {
	user, err := s.repo.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	data := &models.DashboardResponse{}
	if user.UserType == constants.SUPER_ADMIN {
		adminDashboard, err := s.repo.GetSuperAdminDashboard()
		if err != nil {
			logrus.Error("error getting super admin dashboard :: ", err)
			return nil, err
		}
		data.SuperAdmin = adminDashboard
		return data, nil
	}
	if user.UserType == constants.INSTITUTION_ADMIN {
		adminDashboard, err := s.repo.GetInstitutionAdminDashboard(userId)
		if err != nil {
			logrus.Error("error getting institution admin dashboard :: ", err)
			return nil, err
		}
		data.InstitutionAdmin = adminDashboard
		return data, nil
	}
	if user.UserType == constants.PROGRAM_ADMIN {
		adminDashboard, err := s.repo.GetProgramAdminDashboard(userId)
		if err != nil {
			logrus.Error("error getting program admin dashboard :: ", err)
			return nil, err
		}
		data.ProgramAdmin = adminDashboard
		return data, nil
	}
	if user.UserType == constants.TEACHER {
		teacherDashboard, err := s.repo.GetTeacherDashboard(userId)
		if err != nil {
			logrus.Error("error getting teacher dashboard :: ", err)
			return nil, err
		}
		data.Teacher = teacherDashboard
		return data, nil
	}
	if user.UserType == constants.STUDENT {
		studentDashboard, err := s.repo.GetStudentDashboard(userId)
		if err != nil {
			logrus.Error("error getting student dashboard :: ", err)
			return nil, err
		}
		data.Student = studentDashboard
		return data, nil
	}
	return nil, nil
}
