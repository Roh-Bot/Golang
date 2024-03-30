package main

import "fmt"

type MyInterface interface {
	PrintHello()
	PrintDhebug()
}

type MockStruct struct{}

func (m MockStruct) PrintHello() {
	fmt.Println("hello")
}
func (m MockStruct) PrintDhebug() {
	fmt.Println("hello")
}

func main() {
	var m MyInterface
	m = MockStruct{}
	PrintBoth(m)
}

func PrintBoth(m MyInterface) {
	m.PrintHello()
	m.PrintDhebug()
}
