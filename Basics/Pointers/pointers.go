package main

import "fmt"

func main() {
	var p *int
	var number int = 30
	p = &number
	fmt.Println(p, *p, &number, &p)
}
