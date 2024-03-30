package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	//content := "Hello Dhebug"

	//file, err := os.Create("./Hello_Dhebug.txt")
	//checkNilError(err)
	//
	//length, err := file.WriteString(content)
	//checkNilError(err)
	//
	//fmt.Printf("Length is: %d \n", length)

	content1, err := os.ReadFile("Hello_Dhebug.txt")
	checkNilError(err)
	fmt.Println(string(content1))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
