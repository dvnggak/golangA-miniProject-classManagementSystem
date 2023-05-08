package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/constants"
	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/service"
	"github.com/dvnggak/miniProject/utils"
	"github.com/labstack/echo/v4"
)

func (m *Controller) GetUser(c echo.Context) error {

	var users []model.User

	if err := config.DBMysql.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})

}

func (m *Controller) CreateUser(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail",
	}
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	err = service.GetUserRepository().CreateUser(&user)
	if err != nil {
		return err
	}

	data["message"] = "success"
	return c.JSON(http.StatusOK, data)
}

func (m *Controller) LoginUser(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail",
	}
	var loginData struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	err := c.Bind(&loginData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	// Retrieve the user with the given username from the database
	user, err := service.GetUserRepository().GetUserByUsername(loginData.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	// Check if the hashed password and the given password is the same
	if !utils.ComparePassword(user.Password, loginData.Password) {
		return c.JSON(http.StatusBadRequest, data)
	}

	// Generate a JWT token
	token, err := utils.CreateTokenUser(user.ID, user.Username, user.ID_number)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	userResponse := model.UserResponse{Username: user.Username, Message: "Login Succes", Token: token}
	data["message"] = userResponse
	return c.JSON(http.StatusOK, data)
}

func (m *Controller) EnrollClass(c echo.Context) error {
	// Get the class code from the URL parameter
	classCode := c.Param("code")

	// Get the JWT token from the request header
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authorization token is missing",
		})
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid authorization token",
		})
	}

	// Extract the id_number & role field from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid authorization token",
		})
	}
	userID := claims["id_number"].(string)
	role := claims["role"].(string)

	// Check if the user is an admin
	if role != "user" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Only user can enroll in a class",
		})
	}

	// Get the user repository instance
	userRepo := service.GetUserRepository()

	// Get the class by code
	_, err = userRepo.GetClassByCode(classCode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid class code",
		})
	}

	// Check if the user is already enrolled in the class
	if userRepo.CheckEnrolledClass(userID, classCode) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User is already enrolled in this class",
		})
	}

	// Enroll the user in the class
	err = userRepo.EnrollClass(userID, classCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to enroll user in class",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User enrolled in class successfully",
	})
}

func (m *Controller) GetEnrolledClasses(c echo.Context) error {
	userID := c.Param("id_number")

	// Get the JWT token from the request header
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authorization token is missing",
		})
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET_JWT), nil
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid authorization token",
		})
	}

	// Extract the id_number & role field from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid authorization token",
		})
	}
	role := claims["role"].(string)

	// Check if the user is an admin
	if role != "user" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Only user can get enrolled classes",
		})
	}

	userRepo := service.GetUserRepository()
	classes, err := userRepo.GetEnrolledClasses(userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get enrolled classes",
		})
	}

	return c.JSON(http.StatusOK, classes)
}
