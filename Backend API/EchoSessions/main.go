package main

import (
	"github.com/Roh-Bot/EchoSessions/Controllers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	store := sessions.NewCookieStore([]byte("your-secret-key"))
	e := echo.New()
	e.GET("/protected", Controllers.ProtectedHandler, Controllers.RequireLogin)
	e.Use(Controllers.SessionMiddleware(store))

	log.Fatal(e.Start(":8080"))
}
