package controller

import (
	"net/http"

	"github.com/dvnggak/miniProject/config"
	"github.com/dvnggak/miniProject/model"
	"github.com/dvnggak/miniProject/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func (m *Controller) GetUser(c echo.Context) error {
	// user := c.Get("user").(model.User)
	// log.Printf("user data: %+v/n", user)

	var users []model.User

	config.DBMysql.Find(&users)

	return c.JSON(http.StatusOK, users)
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
