package service

import (
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
)

type IUserService interface {
	CreateUser(*model.User) error
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
