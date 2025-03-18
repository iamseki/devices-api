package main

import (
	"github.com/iamseki/devices-api/src/handler"
	"github.com/iamseki/devices-api/src/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	handle := handler.New(repository.New("postgres://devices:devices@localhost:5432/devices?sslmode=disable"))

	e := echo.New()

	e.GET("/devices/:id", handle.GetDevice)
	e.DELETE("/devices/:id", handle.DeleteDevice)

	e.GET("/devices", handle.GetDevice)
	e.POST("/devices", handle.InsertDevice)
	e.PATCH("/devices", handle.UpdateDevice)

	e.Logger.Fatal(e.Start(":8080"))
}
