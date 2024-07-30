package designpatterns

import "fmt"

type EmployeePROT struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *EmployeePROT {
	switch role {
	case Developer:
		return &EmployeePROT{"", "Developer", 60000}
	case Manager:
		return &EmployeePROT{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
