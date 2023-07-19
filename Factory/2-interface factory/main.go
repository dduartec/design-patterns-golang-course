package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	person
}

func (p *person) SayHello() {
	fmt.Printf("Hi my name is %s, I'm %d years old\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
	fmt.Print("I'm too old for this\n")
}

// Factory function
func NewPerson(name string, age int) Person {
	if age > 50 {
		return &tiredPerson{person{name, age}}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("Diego", 22)
	p.SayHello()
	tp := NewPerson("Abel", 68)
	tp.SayHello()
}
