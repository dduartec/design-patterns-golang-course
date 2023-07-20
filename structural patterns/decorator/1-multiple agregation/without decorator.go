package main

import "fmt"

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying")
	}
}

type Lizard struct {
	Age int
}

func (b *Lizard) Crawl() {
	if b.Age >= 10 {
		fmt.Println("Crawling")
	}
}

// Boith embeded structs have age, making difficult its managment
type Dragon struct {
	Bird
	Lizard
}

func (d *Dragon) Age() int {
	return d.Bird.Age
}

// Ugly as fuck
func (d *Dragon) SetAge(age int) {
	d.Bird.Age = 10
	d.Lizard.Age = 10
}

func main() {
	d := Dragon{}
	d.SetAge(11)
	// You can change its values :/
	d.Lizard.Age = 5
	d.Fly()
	d.Crawl()
}
