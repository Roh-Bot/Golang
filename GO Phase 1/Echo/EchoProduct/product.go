package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	products := []map[int]string{{1: "Phone"}, {2: "Laptop"}, {3: "Fridge"}}

	e := echo.New()
	v := validator.New()

	//DEFAULT HOMEPAGE
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Dhebug")
	})

	//GET METHOD TO DISPLAY ALL THE PRODUCTS
	e.GET("/products/:id", func(context echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k, _ := range p {
				pID, err := strconv.Atoi(context.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return context.JSON(http.StatusOK, "Product not found")
		}
		return context.JSON(http.StatusOK, product)
	})
	e.GET("/product", func(context echo.Context) error {
		return context.JSON(http.StatusOK, products)
	})

	//POST METHOD TO APPEND A PRODUCT INTO THE COLLECTION
	e.POST("/products", func(context echo.Context) error {
		type body struct {
			ProdName string `json:"product_name"`
		}
		var requestBody body
		if err := context.Bind(&requestBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: requestBody.ProdName,
		}

		products = append(products, product)
		return context.JSON(http.StatusOK, product)
	})

	//POST METHOD WITH VALIDATION
	e.POST("/products", func(context echo.Context) error {
		type body struct {
			ProdName string `json:"product_name" validate:"required,min=2"`
		}
		var requestBody body

		if err := context.Bind(&requestBody); err != nil {
			return err
		}
		if err := v.Struct(requestBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: requestBody.ProdName,
		}
		products = append(products, product)
		return context.JSON(http.StatusOK, product)
	})

	//PUT METHOD TO UPDATE THE PRODUCTS
	e.PUT("/products", func(context echo.Context) error {
		var product map[int]string
		pID, err := strconv.Atoi(context.Param("id"))
		for _, p := range products {
			for k, _ := range p {
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}

		type body struct {
			ProdName string `json:"product_name" validate:"required,min=2"`
		}
		var requestBody body

		if err := context.Bind(&requestBody); err != nil {
			return err
		}
		if err := v.Struct(requestBody); err != nil {
			return err
		}
		product[pID] = requestBody.ProdName
		return context.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Printf("Listening on Port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
