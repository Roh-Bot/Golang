package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)
	slice1 = append(slice1[:2], slice1[2+1:]...)
	fmt.Println(slice1)
}
