package handlers

import (
	"net/http"

	"github.com/kakuzops/echo-crud/cmd/models"
	"github.com/kakuzops/echo-crud/cmd/repositories"
	"github.com/labstack/echo/v4"
)

func CreateMeasurement(c echo.Context) error {
	meansurement := models.Measurements{}
	c.Bind(&meansurement)
	newMeansurement, err := repositories.CreateMeansurement(meansurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeansurement)
}
