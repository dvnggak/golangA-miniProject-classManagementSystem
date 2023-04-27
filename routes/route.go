package routes

import (
	"github.com/dvnggak/miniProject/constants"
	"github.com/dvnggak/miniProject/controller"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
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

	// restricted group
	eAuth := e.Group("/auth")
	eAuth.Use(JWTMiddleware()) // JWT Middleware

	eAuth.GET("/admins", adminController.GetAdmin)
	// eAuth.GET("/users", userController.GetUser)

	return e
}

func JWTMiddleware() echo.MiddlewareFunc {
	return mid.JWT([]byte(constants.SECRET_JWT))
}
