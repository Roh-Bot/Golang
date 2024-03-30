package Handlers

import (
	"context"
	"fmt"
	"github.com/Roh-Bot/ECHOWITHPGX/Structs"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func RegisterPGX(c echo.Context) error {
	v := validator.New()
	//errValidator := v.RegisterValidation("email_or_mobile", Structs.ValidateEmailOrMobile)

	var conn, err = pgx.Connect(context.Background(), "host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = conn.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	var u Structs.Users
	if err := c.Bind(&u); err != nil {
		return err
	}
	if err := v.Struct(u); err != nil {
		return context.String(500, "Invalid input")
	}

	res := conn.QueryRow(`INSERT INTO "Account".users(display_name, email, mobile, password, mpin) VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		errJson := Structs.ErrorStruct{
			Error: map[string]string{
				"status":  "400",
				"message": "Bad Request",
			},
		}
		return c.JSON(http.StatusBadRequest, errJson)
	} else {
		fmt.Println(res)
		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusOK, u)
	}
	return c.String(http.StatusOK, "ok")
}
