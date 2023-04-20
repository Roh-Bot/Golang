package main

import "fmt"

func main() {
	p1 := Person{"G", 19, "CA"}
	fmt.Println(p1)
	fmt.Println(p1.GetProp())
	fmt.Println(p1)
}

type Person struct {
	Name string
	Age  int
	Job  string
}

func (p Person) GetProp() (string, int, string) { //Struct type receiver
	p.Age = 50
	return p.Name, p.Age, p.Job
}

//type data int
//type data2 string
//
//func (d1 data) multiply(d2 int) data {
//	return data(d2) * d1
//}
//func (d1 data2) multiply(d2 data2) data2 {
//	return d2 + d1
//}
//
//func main() {
//	var num1 data = (2)
//	var result data = num1.multiply(3)
//	fmt.Println("Result:", result)
//
//	var string1 data2 = ("Hello")
//	var string2 data2 = "Dhebug"
//	result2 := string2.multiply(string1)
//	fmt.Println("Result2:", result2)
//}
//var func1 Function = func(s string) string {
//	return s
//}
//fmt.Println(func1("Hello"))
//fmt.Println(func1.Display("Dhebug"))
//}
//
//type Function func(s string) string
//
//func (f Function) Display(s string) string {
//	return s
//}
