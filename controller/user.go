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
