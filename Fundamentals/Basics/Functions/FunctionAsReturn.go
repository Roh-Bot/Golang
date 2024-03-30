package main

import "fmt"

func intSeq() func(s string) int {

	return func(s string) int {
		return 0
	}
}

func FunctionWithNoParameterandIntreturn(s string) int {
	return 0
}

func main() {
	fmt.Println(intSeq())
}
