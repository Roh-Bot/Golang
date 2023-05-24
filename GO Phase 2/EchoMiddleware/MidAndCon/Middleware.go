package MidAndCon

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"os"
)

var e = echo.New()
var v = validator.New()

func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.GET("/", Default)
	e.GET("/products", getAllProducts)
	e.GET("/products/:id", getSingleProduct)
	e.POST("/products", addProduct)
	e.PUT("/products/:id", updateProduct)

	e.Start(fmt.Sprintf("localhost:%s", port))
}
