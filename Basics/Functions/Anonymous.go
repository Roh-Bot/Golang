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

	AnonFuncWithoutCall := func(s string) string {
		return s
	}

	fmt.Println(AnonFuncWithoutCall("Hello Anon"))
}
