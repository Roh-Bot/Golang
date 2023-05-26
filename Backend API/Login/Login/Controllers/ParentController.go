package Controllers

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	e := echo.New()
	e.POST("/register", RegistrationController)
	e.POST("/login", LoginController)
	log.Fatal(e.Start("localhost:8080"))
}
