package designpatterns

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// what if we want factories for specific roles?

// NewEmployeeFactory functional approach
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// EmployeeFactory structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func FactoryGenerator() {
	developer := NewEmployeeFactory("Developer", 60000)("Adam")
	fmt.Println(developer)
	manager := NewEmployeeFactory("Manager", 80000)("Jane")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
