package Controllers

import (
	"github.com/Roh-Bot/StrategyData/StrategyData/Middleware"
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	router := echo.New()
	router.GET("/", Middleware.DefaultPageMiddleware)
	router.POST("strategy/dss", Middleware.DSSMiddleware)
	log.Fatal(router.Start("localhost:8080"))
}
