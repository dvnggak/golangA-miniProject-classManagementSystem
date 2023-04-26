package routes

import (
	"github.com/dvnggak/miniProject/controller"
	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	adminController := controller.Controller{}
	adminGroup := e.Group("/admins")
	adminGroup.POST("/", adminController.CreateAdmin)

	adminGroup.GET("/", adminController.GetAdmin)

	userController := controller.Controller{}
	userGroup := e.Group("/users")
	userGroup.POST("/", userController.CreateUser)

	userGroup.GET("/", userController.GetUser)

	return e
}
