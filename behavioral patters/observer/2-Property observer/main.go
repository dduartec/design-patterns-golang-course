package main

import (
	"container/list"
	"fmt"
)

// Observable

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string // "Age"
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	// Notify on the dependant prop
	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}

	// If we depend on more properties it would be a nightmare!
	// a external structure must be used for this

}

// Cant notify because its a Getter
func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
	o Observable
}

func (t *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congrats, you can vote!")
			t.o.Unsubscribe(t)
		}
	}

}

func main() {
	p := NewPerson(0)
	e := ElectoralRoll{p.Observable}
	p.Subscribe(&e)
	for i := 10; i < 21; i++ {
		fmt.Println("setting age to", i)
		p.SetAge(i)
	}
}
