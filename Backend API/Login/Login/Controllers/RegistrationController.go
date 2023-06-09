package Controllers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/Roh-Bot/Login/Models"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func RegistrationController(c echo.Context) error {
	pool, err := pgxpool.New(context.Background(), "host=localhost port=5432 user=postgres password=admin dbname=User sslmode=disable")
	if err != nil {
		fmt.Println("Connection Failed")
	}

	errPing := pool.Ping(context.Background())
	if errPing != nil {
		log.Fatal(fmt.Println("Connection failed to Databse"))
	} else {
		fmt.Println("DB Connected")
	}

	v := validator.New()

	var response Models.RegistrationResponse
	var register Models.RegistrationStruct

	if err := c.Bind(&register); err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if err := v.Struct(register); err != nil {
		response = Models.RegistrationResponse{
			StatusCode: 400,
			Error: map[string]string{
				"code":    "400",
				"message": "Something went wrong",
			},
			Data: map[string]string{},
		}
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid Input")
	}

	var bytePassword = []byte(register.Password)
	hashedPassword := sha256.Sum256(bytePassword)
	fmt.Println(hashedPassword)

	pgCallStatement := `CALL registration($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	data := []any{register.Email, register.Phone_no, register.Password, register.First_name,
		register.Last_name, register.Dob, register.Address_line_1, register.Address_line_2,
		register.City_id, register.State_id, register.Pincode, register.Referred_by,
		register.Reference_code}

	_, errQuery := pool.Query(context.Background(), pgCallStatement, data...)
	if errQuery != nil {
		fmt.Println(errQuery)
		return c.String(http.StatusInternalServerError, "Query Failed")
	} else {
		fmt.Println("Query Successful")
	}

	response = Models.RegistrationResponse{
		StatusCode: 200,
		Error:      map[string]string{},
		Data: map[string]string{
			"user": "01",
		},
	}
	return c.JSON(http.StatusOK, response)
}
