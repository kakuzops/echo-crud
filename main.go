package main

import (
	storage "github.com/kakuzops/echo-crud/cmd/config"
	"github.com/kakuzops/echo-crud/cmd/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Home)
	// database
	storage.InitDB()
	// --------------
	e.Logger.Fatal(e.Start(":8080"))
}
