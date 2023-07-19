package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

// Factory function
func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("Diego", 22)
	fmt.Println(p)
}
