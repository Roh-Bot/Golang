package Controller

import (
	"github.com/labstack/echo/v4"
	"log"
)

var e = echo.New()

func Start() {
	e.GET("/", Default)
	e.POST("/register", Register)
	log.Fatal(e.Start("localhost:8080"))
}
