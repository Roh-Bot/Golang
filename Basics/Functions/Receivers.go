// Pointer receiver vs Value Receiver in Methods

package main

import "fmt"

type Persons struct {
	Name string
	Age  int
}

// Method with a pointer receiver
func (p *Persons) IncrementAge() {
	p.Age++
}

// Method with a value receiver
func (p Persons) ChangeName(newName string) {
	p.Name = newName
}

func main() {

	p1 := Persons{Name: "Alice", Age: 25}

	// IncrementAge() method modifies the age field of p1
	p1.IncrementAge()
	fmt.Println(p1.Age) // Output: 26

	// ChangeName() method creates a new Persons with the updated name, but doesn't modify p1
	p1.ChangeName("Bob")
	fmt.Println(p1.Name) // Output: Alice
}
