package main

import "fmt"

func PassByValue(x int, y int) { //Makes a copy of variables a,b
	temp := x
	x = y
	y = temp

}

func PassByReference(x *int, y *int) { //Passes the memory address which holds the value of a & b
	temp := *x
	*x = *y
	*y = temp

}

func main() {
	a := 10
	b := 20
	PassByValue(a, b)
	fmt.Printf("a:%d, b:%d\n", a, b) //Passing by value -- Makes a copy of a,b

	PassByReference(&a, &b)
	fmt.Printf("a:%d, b:%d", a, b) //Passing the adress which holds the actual value

}
