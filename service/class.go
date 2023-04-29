package service

import (
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
)

type IClassService interface {
	CreateClass(*model.Class) error
	UpdateClass(*model.Class) error
	DeleteClass(*model.Class) error
}

type ClassRepository struct {
	Func IClassService
}

var classRepository IClassService

func init() {
	ur := &ClassRepository{}
	ur.Func = ur

	classRepository = ur
}

func GetClassRepository() IClassService {
	return classRepository
}

func SetClassRepository(ur IClassService) {
	classRepository = ur
}

func (u *ClassRepository) CreateClass(class *model.Class) error {
	err := config.DBMysql.Save(&class)
	if err != nil {
		return err.Error
	}

	return nil
}

func (u *ClassRepository) UpdateClass(class *model.Class) error {
	err := config.DBMysql.Save(&class)
	if err != nil {
		return err.Error
	}

	return nil
}

func (u *ClassRepository) DeleteClass(class *model.Class) error {
	err := config.DBMysql.Delete(&class)
	if err != nil {
		return err.Error
	}

	return nil
}
