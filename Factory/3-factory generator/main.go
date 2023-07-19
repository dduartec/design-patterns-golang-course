package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

// functional approach
func NewEmployeeFactoryFunc(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, income}
	}
}

// struct approach
type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Income}
}

func NewEmployeeFactory(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{position, income}
}

func main() {
	developerFactory := NewEmployeeFactoryFunc("dev", 6000)
	managerFactory := NewEmployeeFactoryFunc("manager", 8000)

	dev := developerFactory("Diego")
	manager := managerFactory("Abel")

	fmt.Println(dev)
	fmt.Println(manager)

	bossfactory := NewEmployeeFactory("CEO", 10000)
	boss := bossfactory.Create("Sam")
	fmt.Println(boss)
}
