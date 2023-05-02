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
	CheckEnrolledClass(string, string) bool
	EnrollClass(string, string) error
	GetEnrolledClasses(string) ([]model.Class, error)
	UnenrollUserFromClass(string, string) error
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
func (r *UserRepository) CheckEnrolledClass(id_number string, classCode string) bool {
	user := &model.User{}
	result := config.DBMysql.Preload("Enrolled", "code = ?", classCode).Where("id_number = ?", id_number).First(user)
	if result.Error != nil {
		return false
	}
	return len(user.Enrolled) > 0
}

// EnrollClass enrolls the user with the given ID to the class with the given ID
func (r *UserRepository) EnrollClass(id_number string, classCode string) error {
	// Get the user by ID number
	var user model.User
	if err := config.DBMysql.Where("id_number = ?", id_number).First(&user).Error; err != nil {
		return err
	}

	// Get the class by code
	var class model.Class
	if err := config.DBMysql.Where("code = ?", classCode).Preload("Enrolled").First(&class).Error; err != nil {
		return err
	}

	// Check if the user is already enrolled in the class
	for _, c := range user.Enrolled {
		if c.Code == classCode {
			return errors.New("user is already enrolled in this class")
		}
	}

	// Append the class to the user's enrolled classes
	user.Enrolled = append(user.Enrolled, class)

	// Append the user to the class's enrolled users
	class.Enrolled = append(class.Enrolled, &user)

	// Save the changes to the database
	if err := config.DBMysql.Save(&user).Error; err != nil {
		return err
	}
	if err := config.DBMysql.Save(&class).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UnenrollUserFromClass(classCode string, userID string) error {
	// Get the user by ID number
	var user model.User
	if err := config.DBMysql.Where("id_number = ?", userID).First(&user).Error; err != nil {
		return err
	}

	// Get the class by code
	var class model.Class
	if err := config.DBMysql.Where("code = ?", classCode).Preload("Enrolled").First(&class).Error; err != nil {
		return err
	}
	// Check if the user is enrolled in the class
	if !r.CheckEnrolledClass(userID, classCode) {
		return errors.New("user is not enrolled in this class")
	}

	// Remove the user from the class's enrolled users
	for i, u := range class.Enrolled {
		if u.ID_number == user.ID_number {
			class.Enrolled = append(class.Enrolled[:i], class.Enrolled[i+1:]...)
			break
		}
	}

	// Save the changes to the database
	if err := config.DBMysql.Save(&user).Error; err != nil {
		return err
	}
	if err := config.DBMysql.Save(&class).Error; err != nil {
		return err
	}

	return nil

}
func (r *UserRepository) GetEnrolledClasses(userID string) ([]model.Class, error) {
	var user model.User
	if err := config.DBMysql.Preload("Enrolled").First(&user, "id_number = ?", userID).Error; err != nil {
		return nil, err
	}

	return user.Enrolled, nil
}
