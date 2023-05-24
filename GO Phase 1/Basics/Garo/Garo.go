package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x, y, z = 1, 2, 3
	fmt.Printf("Hello "+
		"world %c %d %v ", x, y, z)

	fmt.Println(`
				Garo vei
				Bottom Text
				`)
	var rune1 int32 = 'A'
	fmt.Printf("%d \n", rune1)
	fmt.Println(reflect.TypeOf(rune1))

}
