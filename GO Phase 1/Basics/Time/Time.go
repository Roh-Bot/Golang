package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome")

	//fmt.Println(time.Now())
	//
	fmt.Println(time.Now().Format("01-02-2006 15:04:05 2006"))

	fmt.Println(time.Date(2001, time.May, 12, 12, 30, 30, 00, time.Local))
}
