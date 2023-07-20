package main

import "fmt"

// decorator interface
type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	age int
}

func (b *Lizard) Age() int {
	return b.age
}

func (b *Lizard) SetAge(age int) {
	b.age = age
}

func (b *Lizard) Crawl() {
	if b.age >= 10 {
		fmt.Println("Crawling")
	}
}

// make embeded structs private
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (b *Dragon) Age() int {
	return b.bird.age
}

func (b *Dragon) SetAge(age int) {
	b.lizard.age = age
	b.bird.age = age
}

func (b *Dragon) Fly() {
	b.bird.Fly()
}

func (b *Dragon) Crawl() {
	b.lizard.Crawl()
}

func main() {
	d := Dragon{}
	d.SetAge(11)
	// Now the user cant access the values and have to use the methods
	// d.Lizard.Age = 5
	d.Fly()
	d.Crawl()
}
