package main

import (
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/routes"
)

func init() {
	// godotenv.Load(".env")
	config.InitDB()
	model.InitMigrate()
}

func main() {
	route := routes.StartRoute()

	route.Start(":8080")
}
