package main

import "fmt"

func main() {
	func(s string) string {
		return s
	}("Hello World")

	AnonFunc := func(s string) string {
		return s
	}("Hello Anon")

	fmt.Println(AnonFunc)

	AnonFuncWithoutParam := func(s string) string {
		return s
	}

	fmt.Println(AnonFuncWithoutParam("Hello Anon"))
}
