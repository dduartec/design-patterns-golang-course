package main

import "fmt"

// Dependency Inversion Principle
// High Level Modules should not depend on Low Level Modules
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from        *Person
	relatioship Relationship
	to          *Person
}

// low level model (storage mechanism)
type Relationships struct {
	relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high level model (operate over data)
type Research struct {
	// break DIP
	relationships *Relationships
}

func (r *Research) Investigate() {
	relationships := r.relationships.relations
	for _, rel := range relationships {
		if rel.relatioship == Parent {
			fmt.Printf("%s has a child called %s\n", rel.from.name, rel.to.name)
		}
	}
}

// DIP

type RelationshipBrowser interface {
	FindAllFrom(r Relationship, from string) []*Person
}

type BetterRelationships []Info

func (r *BetterRelationships) AddParentAndChild(parent, child *Person) {
	*r = append(*r, Info{parent, Parent, child})
	*r = append(*r, Info{child, Child, parent})
}

func (br BetterRelationships) FindAllFrom(r Relationship, from string) []*Person {
	relationships := br
	result := make([]*Person, 0)
	for _, rel := range relationships {
		if rel.relatioship == r && rel.from.name == from {
			result = append(result, rel.to)
		}
	}
	return result
}

// Using an interface we donot deppend of the implementation (low level details)
type BetterResearch struct {
	browser RelationshipBrowser
}

func (r *BetterResearch) Investigate(from string) {
	res := r.browser.FindAllFrom(Parent, from)
	for _, v := range res {
		fmt.Printf("%s has a child called %s\n", from, v.name)
	}
}

func main() {
	parent := Person{
		name: "Jhon",
	}
	c1 := Person{
		name: "Chris",
	}
	c2 := Person{
		name: "Matt",
	}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &c1)
	relationships.AddParentAndChild(&parent, &c2)

	r := Research{&relationships}
	r.Investigate()

	fmt.Println("---DIP---")

	betterRelationships := BetterRelationships{}
	betterRelationships.AddParentAndChild(&parent, &c1)
	betterRelationships.AddParentAndChild(&parent, &c2)

	br := BetterResearch{betterRelationships}
	br.Investigate("Jhon")
}
