package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
	Iterator                        *PersonNameIterator
}

func NewPerson(first, middle, last string) *Person {
	p := &Person{FirstName: first, MiddleName: middle, LastName: last}
	it := NewPersonNameIterator(p)
	p.Iterator = it
	return p
}

func (p *Person) Names() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGen() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()

	return out
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("ERROR")
}

func main() {
	p := NewPerson("Diego", "Duarte", "Camacho")
	p1 := NewPerson("Diego", "", "Camacho")

	for _, name := range p.Names() {
		fmt.Println(name)
	}

	for name := range p1.NamesGen() {
		fmt.Println(name)
	}

	for it := p.Iterator; it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
