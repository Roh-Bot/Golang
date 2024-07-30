package designpatterns

import "fmt"

type calc interface {
	area() int
}

type rect struct {
	width, height int
}

type square struct {
	height int
}

type Circle struct {
	Radius int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r square) area() int {
	return r.height * r.height
}

func SimpleInterfaceExecute() {
	var c []calc
	c = []calc{square{height: 10}, rect{width: 10, height: 5}}
	for _, v := range c {
		fmt.Println(v.area())
	}
}
