package model

import (
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/utils"
	"gorm.io/gorm"
)

func InitMigrate() {
	config.DBMysql.AutoMigrate(&User{})
	config.DBMysql.AutoMigrate(&Admin{})
	config.DBMysql.AutoMigrate(&Class{})
}

type Class struct {
	gorm.Model
	Code          string  `gorm:"primaryKey" json:"code" form:"code"`
	Course_name   string  `json:"course_name" form:"course_name"`
	Lecturer_name string  `json:"lecturer_name" form:"lecturer_name"`
	Class_type    string  `json:"class_type" form:"class_type"`
	Day           string  `json:"day" form:"day"`
	Classroom     string  `json:"classroom" form:"classroom"`
	Period        string  `json:"period" form:"period"`
	Link_group    string  `json:"link_group" form:"link_group"`
	Enrolled      []*User `gorm:"many2many:user_classes;"`
}

type Admin struct {
	gorm.Model
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type AdminResponse struct {
	Username string `json:"name" form:"name"`
	Message  string `json:"message" form:"message"`
	Token    string `json:"token" form:"token"`
}

type User struct {
	gorm.Model
	ID_number string  `form:"id_number" json:"id_number"`
	Name      string  `form:"name" json:"name"`
	Type      string  `form:"type" json:"type"`
	Email     string  `form:"email" json:"email"`
	Username  string  `form:"username" json:"username"`
	Password  string  `form:"password" json:"password"`
	Enrolled  []Class `gorm:"many2many:user_classes;"`
}

type UserResponse struct {
	Username string `json:"username" form:"username"`
	Message  string `json:"message" form:"message"`
	Token    string `json:"token" form:"token"`
}

func (u *User) BeforeCreateUser(tx *gorm.DB) (err error) {
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword

	return
}

func (a *Admin) BeforeCreateAdmin(tx *gorm.DB) (err error) {
	hashPassword, err := utils.HashPassword(a.Password)
	if err != nil {
		return err
	}
	a.Password = hashPassword

	return
}

func (a *Admin) ComparePassword(password string) string {
	utils.ComparePassword(a.Password, password)
	return a.Password
}
