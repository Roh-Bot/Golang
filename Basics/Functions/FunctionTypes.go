package main

import "fmt"

type data int
type data2 string

// Data
func (d1 data) multiply(d2 int) data {
	return data(d2) * d1
}
func (d1 data2) multiply(d2 data2) data2 {
	return d2 + d1
}

type FunctionType func(s string) string

func (f FunctionType) Display(s string) string {
	return s
}

func main() {
	var num1 data = (2)
	var result data = num1.multiply(3)
	fmt.Println("Result:", result)
	var string1 data2 = ("Hello")
	var string2 data2 = "Dhebug"
	result2 := string2.multiply(string1)
	fmt.Println("Result2:", result2)

	var func1 FunctionType = func(s string) string {
		return s
	}
	fmt.Println(func1("Hello"))
	fmt.Println(func1.Display("Dhebug"))
	//fmt.Println(FunctionType(func(s string) string {
	//	return s
	//}))
}
