package main

import "fmt"

func main() {
	func(s string) string {
		return s
	}("Hello World")

	var AnonFunc string
	AnonFunc = func(s string) string {
		return s
	}("Hello Anon")

	AnonFunc = myFunc("Hello Anon")

	fmt.Println(AnonFunc)

	//var AnonFuncWithoutCall func(s string) string
	var AnonFuncWithoutCall func(s string) string

	AnonFuncWithoutCall = func(s string) string {
		myString := "Hello"
		return myString
	}

	fmt.Println(AnonFuncWithoutCall("Hello Anon"))
	fmt.Println(AnonFunc)
}
func myFunc(s string) string {
	return s
}
