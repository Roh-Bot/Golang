package db

import (
	"log"
)

func init() {
	err := postgresNewPool()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
