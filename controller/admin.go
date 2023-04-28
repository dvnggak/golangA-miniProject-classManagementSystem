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
		"message": "fail to create admin",
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

	data["message"] = "success to create admin"
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

func (m *Controller) CreateClass(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail to create class",
	}

	var class model.Class
	err := c.Bind(&class)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}

	err = service.GetClassRepository().CreateClass(&class)
	if err != nil {
		return err
	}

	data["message"] = "success to create class"
	return c.JSON(http.StatusOK, data)
}

func (m *Controller) GetClass(c echo.Context) error {
	var classes []model.Class

	if err := config.DBMysql.Find(&classes).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all classes",
		"classes": classes,
	})
}

func (m *Controller) UpdateClass(c echo.Context) error {
	code := c.Param("code")
	var class model.Class

	if err := config.DBMysql.Where("code = ?", code).First(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&class); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := service.GetClassRepository().UpdateClass(&class)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update class",
		"class":   class,
	})
}

func (m *Controller) DeleteClass(c echo.Context) error {
	code := c.Param("code")
	var class model.Class

	if err := config.DBMysql.Where("code = ?", code).First(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := service.GetClassRepository().DeleteClass(&class)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete class",
		"class":   class,
	})
}
