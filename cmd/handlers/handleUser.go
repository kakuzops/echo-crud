package handlers

import (
	"net/http"

	"github.com/kakuzops/echo-crud/cmd/models"
	"github.com/kakuzops/echo-crud/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	newUser, err := repositories.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}
