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
	GetClassByCode(string) (*model.Class, error)
	CheckEnrolledClass(string, uint) bool
	EnrollClass(string, uint) error
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
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetClassByCode retrieves the class with the given code from the database
func (r *UserRepository) GetClassByCode(code string) (*model.Class, error) {
	var class model.Class
	result := config.DBMysql.Where("code = ?", code).First(&class)
	if result.Error != nil {
		return nil, result.Error
	}
	return &class, nil
}

// CheckEnrolledClass checks if the user with the given ID is already enrolled in the class with the given ID
func (r *UserRepository) CheckEnrolledClass(id_number string, classID uint) bool {
	var user model.User
	var class model.Class

	config.DBMysql.Model(&user).Where("id_number = ?", id_number).Preload("Enrolled", "id = ?", classID).Find(&user)
	config.DBMysql.Model(&class).Where("id = ?", classID).Preload("Enrolled", "id_number = ?", id_number).Find(&class)

	return len(user.Enrolled) > 0 && len(class.Enrolled) > 0
}

// EnrollClass enrolls the user with the given ID to the class with the given ID
func (r *UserRepository) EnrollClass(id_number string, classID uint) error {
	var user model.User
	var class model.Class

	if err := config.DBMysql.Where("id_number = ?", id_number).First(&user).Error; err != nil {
		return err
	}
	if err := config.DBMysql.Where("id = ?", classID).First(&class).Error; err != nil {
		return err
	}

	err := config.DBMysql.Model(&user).Association("Enrolled").Append(&class)
	if err != nil {
		return err
	}
	return nil
}
