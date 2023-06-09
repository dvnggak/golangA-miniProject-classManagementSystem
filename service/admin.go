package service

import (
	"errors"

	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"gorm.io/gorm"
)

type IAdminService interface {
	CreateAdmin(*model.Admin) error
	GetAdminByUsername(string) (*model.Admin, error)
	GetEnrolledUsersByClassCode(string) ([]*model.User, error)
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

func (u *AdminRepository) GetAdminByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	if err := config.DBMysql.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return &admin, nil
}

func (r *AdminRepository) GetEnrolledUsersByClassCode(code string) ([]*model.User, error) {
	var users []*model.User
	err := config.DBMysql.Joins("JOIN user_classes ON user_classes.user_id = users.id").
		Where("user_classes.class_code = ?", code).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
