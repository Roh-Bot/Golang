package main

import "fmt"

type intData int
type data2 string

// Data
func (d1 intData) multiply(d2 int) intData {
	return intData(d2) * d1
}
func (d1 data2) multiply(d2 data2) data2 {
	return d2 + d1
}

//func (d1 int) multiply2(d2 int) int{
//	return 0
//}

type FunctionType func(s string) string

func (f FunctionType) Display() FunctionType {
	return func(s string) string {
		return "Hello"
	}
}

func main() {
	var intDataObject intData
	fmt.Println(intDataObject.multiply(3))
	//var string1 data2 = ("Hello")
	//var string2 data2 = "Dhebug"
	//result2 := string2.multiply(string1)
	//fmt.Println("Result2:", result2)

	var func1 FunctionType = func(s string) string {
		return s
	}
	fmt.Println(func1("Hello"))
	fmt.Println(func1.Display())
	//fmt.Println(FunctionType(func(s string) string {
	//	return s
	//}))
}
