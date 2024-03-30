package main

import "fmt"

func main() {
	var slice1 = []int{0, 0, 0}
	slice2 := make([]int, 4, 4)
	fmt.Println(len(slice2))
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)
	slice1 = append(slice1[:2], slice1[2+1:]...)
	fmt.Println(slice1)
}
