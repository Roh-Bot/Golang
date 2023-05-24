package main

import (
	"fmt"
	"log"

	"github.com/upper/db/v4/adapter/postgresql"
)

var settings=postgresql.ConnectionURL{
	Host: "localhost",
	Database: "acount_profile",
	User: "postgres",
	Password: "root",

}

type User struct{
	ID int `db:"id,omitempty"`
	FirstName string `db:"first_name"`
	LastName string  `db:"last_name"`
}

func main() {
	sess,err:=postgresql.Open(settings)
	chexkError(err)

	// tables,err:=sess.Collections()
	// chexkError(err)

	// testTb:= sess.Collection("test")

	// res,err:=testTb.Insert(User{
	// 	FirstName: "Atul",
	// 	LastName: "Pandey",
	// })
	// chexkError(err)

	// fmt.Printf("%v",res.ID())
	testTb:= sess.Collection("test")
	res,err:=testTb.Insert(map[string]string{
		"first_name": "Atul",
		"last_name": "Pandey",
	})
	chexkError(err)

	fmt.Printf("%v",res.ID())


}

func chexkError(err error){
	if err!= nil{
		log.Fatal(err)
	}


}