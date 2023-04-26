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
	config.DBMysql.AutoMigrate(&ClassDetail{})
}

type Class struct {
	gorm.Model
	Code          string `gorm:"primaryKey" json:"code" form:"code"`
	Course_name   string `json:"course_name" form:"course_name"`
	Lecturer_name string `json:"lecturer_name" form:"lecturer_name"`
	Class_type    string `json:"class_type" form:"class_type"`
	Day           string `json:"day" form:"day"`
	Classroom     string `json:"classroom" form:"classroom"`
	Period        string `json:"period" form:"period"`
	Link_group    string `json:"link_group" form:"link_group"`
}

type ClassDetail struct {
	gorm.Model
	ClassCode string `json:"class_code" form:"class_code"` // foreign key reference to Class Code
	Users     []User `gorm:"many2many:class_details;"`
}

type Admin struct {
	gorm.Model
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type User struct {
	gorm.Model
	ID_number    int           `form:"id_number" json:"id_number"`
	Name         string        `form:"name" json:"name"`
	Type         string        `form:"type" json:"type"`
	Email        string        `form:"email" json:"email"`
	Username     string        `form:"username" json:"username"`
	Password     string        `form:"password" json:"password"`
	ClassDetails []ClassDetail `gorm:"many2many:user_classes;"`
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
