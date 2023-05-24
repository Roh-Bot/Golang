package MidAndCon

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var products = []map[int]string{{1: "Phone"}, {2: "Laptop"}, {3: "Fridge"}}

func Default(context echo.Context) error {
	return context.String(http.StatusOK, "Hello Dhebug")
}

func getSingleProduct(context echo.Context) error {
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
}
func getAllProducts(context echo.Context) error {
	return context.JSON(http.StatusOK, products)
}

func addProduct(context echo.Context) error {
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
}

func updateProduct(context echo.Context) error {
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
}

func deleteProduct(context echo.Context) error {
	var product map[int]string
	var index int
	pId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return context.String(404, "Invalid id")
	}
	for i, p := range products {
		for k, _ := range p {
			if pId == k {
				product = p
				index = i
			}
		}
	}

	if product != nil {
		return context.String(404, "Product Not Found")
	}

	splice := func(products []map[int]string, i int) []map[int]string {
		return append(products)
	}
}
