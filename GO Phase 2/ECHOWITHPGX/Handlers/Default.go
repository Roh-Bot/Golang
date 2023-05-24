package Handlers

import "github.com/labstack/echo/v4"

func Default(context echo.Context) error {
	return context.String(200, "Default Page")
}
