package Controller

import (
	"database/sql"
	"echodb/Structs"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func Default(context echo.Context) error {
	return context.String(http.StatusOK, "Default Page")
}

func Register(context echo.Context) error {
	v := validator.New()
	var db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	var u Structs.User
	if err := context.Bind(&u); err != nil {
		return err
	}
	if err := v.Struct(u); err != nil {
		return context.String(500, "Name Required")
	}

	sqlStatement := `INSERT INTO "Users" (email, "fName", "lName","phone") VALUES ($1, $2, $3, $4)`
	res, err := db.Query(sqlStatement, u.Email, u.Fname, u.Lname, u.Phone)
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
