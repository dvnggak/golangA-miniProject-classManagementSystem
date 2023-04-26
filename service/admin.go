package service

import (
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
)

type IAdminService interface {
	CreateAdmin(*model.Admin) error
}

type AdminRepository struct {
	Func IAdminService
}

var adminRepository IAdminService

func init() {
	ur := &AdminRepository{}
	ur.Func = ur

	adminRepository = ur
}

func GetAdminRepository() IAdminService {
	return adminRepository
}

func SetAdminRepository(ur IAdminService) {
	adminRepository = ur
}

func (u *AdminRepository) CreateAdmin(admin *model.Admin) error {
	admin.BeforeCreateAdmin(config.DBMysql)
	err := config.DBMysql.Save(&admin)
	if err != nil {
		return err.Error
	}

	return nil
}
