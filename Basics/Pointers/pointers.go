package main

import "fmt"

func main() {
	var p *int
	//p := &number
	var number int = 30

	fmt.Println(p, *p, &number, &p)
}

//type User struct {
//	Name string `example:"name"`
//}
//
//func (u *User) String() string {
//	return fmt.Sprintf("Hi! My name is %s", u.Name)
//}
//
//func main() {
//	u := &User{
//		Name: "Sammy",
//	}
//
//	fmt.Println(u)
//}
