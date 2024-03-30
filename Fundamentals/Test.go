package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	fmt.Println(os.Environ())
	pool := PostgresConnectionPool()
	_, errx := pool.Query(context.Background(), `Select * from users`)
	if errx != nil {
		log.Fatal("", errx)
	}
}

func PostgresConnectionPool() *pgxpool.Pool {
	conf, err := pgxpool.ParseConfig("")
	if err != nil {
		fmt.Println("Connection Failed")
	}
	fmt.Println(conf)
	var config pgx.ConnConfig
	fmt.Println(config.ConnString())

	pool, errPing := pgxpool.NewWithConfig(context.Background(), conf)
	if errPing != nil {
		log.Fatal(fmt.Println("Connection failed to Databse"))
	} else {
		fmt.Println("DB Connected")
	}
	return pool
}
