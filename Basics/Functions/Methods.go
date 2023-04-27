package main

import "fmt"

func main() {
	p1 := Person{"G", 19, "CA"}
	fmt.Println(p1)
	fmt.Println(p1.GetProp())
	fmt.Println(p1)
}

type Person struct {
	Name string
	Age  int
	Job  string
}

func (p Person) GetProp() (string, int, string) { //Struct type receiver
	p.Age = 50
	return p.Name, p.Age, p.Job
}
