package designpatterns

import "fmt"

type PersonF struct {
	Name string
	Age  int
}

// factory function
//func NewPersonF(name string, age int) PersonF {
//  return PersonF{name, age}
//}

func NewPersonF(name string, age int) *PersonF {
	return &PersonF{name, age}
}

func FactoryFunction() {
	// initialize directly
	p := PersonF{"John", 22}
	fmt.Println(p)

	// use a constructor
	p2 := NewPersonF("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
