package controller

import (
	"net/http"

	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/service"
	"github.com/dvnggak/miniProject/utils"
	"github.com/labstack/echo/v4"
)

func (m *Controller) GetAdmin(c echo.Context) error {
	var admins []model.Admin

	if err := config.DBMysql.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"users":   admins,
	})
}

func (m *Controller) CreateAdmin(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail",
	}
	var admin model.Admin
	err := c.Bind(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	err = service.GetAdminRepository().CreateAdmin(&admin)
	if err != nil {
		return err
	}

	data["message"] = "success"
	return c.JSON(http.StatusOK, data)
}

func (m *Controller) LoginAdmin(c echo.Context) error {
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

	// Retrieve the admin with the given username from the database
	admin, err := service.GetAdminRepository().GetAdminByUsername(loginData.Username)
	if err != nil {
		return err
	}

	// Verify the password
	if !utils.ComparePassword(admin.Password, loginData.Password) {
		data["message"] = "incorrect password"
		return c.JSON(http.StatusUnauthorized, data)
	}

	token, err := utils.CreateTokenAdmin(admin.ID, admin.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "failed to create token",
			"error":   err.Error(),
		})
	}

	adminResponse := model.AdminResponse{Username: admin.Username, Message: "Login success", Token: token}

	data["message"] = adminResponse
	return c.JSON(http.StatusOK, data)
}
