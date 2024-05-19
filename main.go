package main

import (
	storage "github.com/kakuzops/echo-crud/cmd/config"
	"github.com/kakuzops/echo-crud/cmd/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	// database
	storage.InitDB()
	// --------------

	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
