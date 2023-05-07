package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBMysql *gorm.DB

func InitDB() {
	var err error

	dbUsername := "root"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "class_management_system"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("connection mysql:", connectionString)
	DBMysql, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
