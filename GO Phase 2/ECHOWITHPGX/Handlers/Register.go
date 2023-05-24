package Handlers

import (
	"database/sql"
	"fmt"
	"github.com/Roh-Bot/ECHOWITHPGX/Structs"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Register(context echo.Context) error {
	v := validator.New()
	//errValidator := v.RegisterValidation("email_or_mobile", Structs.ValidateEmailOrMobile)

	var db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	var u Structs.Users
	if err := context.Bind(&u); err != nil {
		return err
	}
	if err := v.Struct(u); err != nil {
		return context.String(500, "Invalid input")
	}

	sqlStatement := `INSERT INTO "Account".users(display_name, email, mobile, password, mpin) VALUES ($1, $2, $3, $4, $5)`
	res, err := db.Query(sqlStatement, u.Display_name, u.Email, u.Mobile, u.Password, u.Mpin)
	if err != nil {
		errJson := Structs.ErrorStruct{
			Error: map[string]string{
				"status":  "400",
				"message": "Bad Request",
			},
		}
		return context.JSON(http.StatusBadRequest, errJson)
	} else {
		fmt.Println(res)
		context.Response().Header().Set("Content-Type", "application/json")
		return context.JSON(http.StatusOK, u)
	}
	return context.String(http.StatusOK, "ok")
}
