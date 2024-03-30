package Controllers

import (
	"github.com/Roh-Bot/ECHOWITHPGX/Handlers"
	"github.com/labstack/echo/v4"
	"log"
)

var e = echo.New()

func Start() {
	e.GET("/", Handlers.Default)
	e.POST("/register", Handlers.Register)
	e.POST("/registerpgx", Handlers.RegisterPGX)
	log.Fatal(e.Start("localhost:8080"))
}
