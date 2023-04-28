package service

import (
	"errors"

	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"gorm.io/gorm"
)

type IUserService interface {
	CreateUser(*model.User) error
	GetUserByUsername(string) (*model.User, error)
}

type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() IUserService {
	return userRepository
}

func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) CreateUser(user *model.User) error {
	user.BeforeCreateUser(config.DBMysql)
	err := config.DBMysql.Save(&user)
	if err != nil {
		return err.Error
	}

	return nil
}

func (u *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := config.DBMysql.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("admin not found")
		}
		return nil, err
	}
	return &user, nil
}
