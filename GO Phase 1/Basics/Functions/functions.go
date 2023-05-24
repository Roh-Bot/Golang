package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")

	fmt.Println(proadder([]int{1, 2, 3, 4}...))

	fmt.Println(proadder(1, 2, 3, 4))

}

func adder(val1, val2 int) int {
	return val1 + val2
}

func proadder(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "Hello Dhebug"
}
