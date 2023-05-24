package main

import "fmt"

//func main() {
//	var p *int
//	var number int = 30
//	p = &number
//
//	fmt.Println(p, *p, &number, &p)
//}

type User struct {
	Name string
}

func (u1 *User) String() string {
	u1.Name = "Garo"
	return u1.Name
}

func main() {

	u := &User{
		Name: "Sammy",
	}

	fmt.Println("Function print", u.String())
	fmt.Println("main print", u.Name)
}
