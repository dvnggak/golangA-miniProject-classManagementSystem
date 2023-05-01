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
	userGroup.POST("/login", userController.LoginUser)

	// restricted group
	eAuth := e.Group("/auth")
	eAuth.Use(JWTMiddleware()) // JWT Middleware

	// Admin routes
	eAuth.GET("/admins", adminController.GetAdmin)

	// Class routes
	eAuth.POST("/admins/newClass", adminController.CreateClass)
	eAuth.GET("/admins/classes", adminController.GetClass)
	eAuth.PUT("/admins/classes/:code", adminController.UpdateClass)
	eAuth.DELETE("/admins/classes/:code", adminController.DeleteClass)
	eAuth.GET("/admins/classes/:code/users", adminController.GetEnrolledUsers)

	// User routes
	eAuth.GET("/users", userController.GetUser)
	// User enroll class
	eAuth.POST("/users/enroll/:code", userController.EnrollClass)
	// User get enrolled class
	eAuth.GET("/users/:id_number/classes", userController.GetEnrolledClasses)

	return e
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := mid.JWTConfig{
		SigningKey: []byte(constants.SECRET_JWT),
	}
	return mid.JWTWithConfig(config)
}
