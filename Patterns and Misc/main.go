package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 1
	fmt.Println(reflect.TypeFor[int]())
	fmt.Println(reflect.TypeOf(num))
}
