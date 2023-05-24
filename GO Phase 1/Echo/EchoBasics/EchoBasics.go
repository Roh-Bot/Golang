package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Dhebug")
	})
	e.GET("/:vendor", func(c echo.Context) error {
		return c.String(200, c.Param("vendor"))
	})
	e.GET("/:vendor?model=iPhone14", func(c echo.Context) error {
		return c.String(200, c.Param("vendor"))
	})
	e.Logger.Fatal(e.Start("localhost:8080"))
}
