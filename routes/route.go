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
	adminGroup.POST("/login", adminController.LoginAdmin)

	userController := controller.Controller{}
	userGroup := e.Group("/users")
	userGroup.POST("/", userController.CreateUser)

	adminGroup.GET("/", adminController.GetAdmin)
	userGroup.GET("/", userController.GetUser)

	return e
}
