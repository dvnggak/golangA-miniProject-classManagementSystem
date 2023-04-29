package controller

import (
	"net/http"

	"github.com/dvnggak/miniProject/config"
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
	token, err := utils.CreateTokenUser(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	userResponse := model.UserResponse{Username: user.Username, Message: "Login Succes", Token: token}
	data["message"] = userResponse
	return c.JSON(http.StatusOK, data)
}

func (m *Controller) EnrollClass(c echo.Context) error {
	// // fmt.Printf("Request Headers: %v\n", c.Request().Header)

	data := map[string]interface{}{
		"message": "fail to enroll class",
	}
	var enrollData struct {
		Code         string `json:"code" validate:"required"`
		UserIDNumber string `json:"id_number" validate:"required"`
	}
	err := c.Bind(&enrollData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	data["message"] = "process to enroll class"
	// Get the class with the given code
	class, err := service.GetUserRepository().GetClassByCode(enrollData.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	data["message"] = "checkEnroll "
	// Check if the user with the given ID is already enrolled in the class
	if service.GetUserRepository().CheckEnrolledClass(enrollData.UserIDNumber, class.ID) {
		return c.JSON(http.StatusBadRequest, data)
	}

	// Enroll the user with the given ID to the class with the given ID
	if err := service.GetUserRepository().EnrollClass(enrollData.UserIDNumber, class.ID); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	data["message"] = "success"
	return c.JSON(http.StatusOK, data)
}
