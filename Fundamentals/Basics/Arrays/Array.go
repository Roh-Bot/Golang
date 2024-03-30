package main

import (
	"fmt"
)

func main() {
	var arr1 []int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1, len(arr1))
	fmt.Println(arr2)
	fmt.Printf("%T", arr1)
}
