package main

import "fmt"

func main() {
	defer fmt.Println("Defer1")
	defer fmt.Println("Defer2")
	defer fmt.Println("Defer3")
	fmt.Println("Hello world")
}
