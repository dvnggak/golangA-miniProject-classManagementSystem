package controller

import (
	"net/http"

	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/service"
	"github.com/labstack/echo/v4"
)

func (m *Controller) GetAdmin(c echo.Context) error {
	var admins []model.Admin

	config.DBMysql.Find(&admins)

	return c.JSON(http.StatusOK, admins)
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
