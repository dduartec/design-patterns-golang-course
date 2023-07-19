package main

import "fmt"

type Employee struct {
	Name, Position string
	Income         int
}

// prototypes
type Role int

const (
	Developer Role = iota
	Manager
)

func NewEmployee(name string, role Role) *Employee {
	switch role {
	case Developer:
		return &Employee{name, "dev", 6000}
	case Manager:
		return &Employee{name, "manager", 6000}
	default:
		panic("unsupported role")
	}
}

func main() {
	dev := NewEmployee("Diego", Developer)
	manager := NewEmployee("Abel", Manager)

	fmt.Println(dev)
	fmt.Println(manager)
}
