package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{name, attack, defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c *CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(creature *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: creature}}
}

func (c *DoubleAttackModifier) Handle() {
	fmt.Printf("Doubling creature's %s attack", c.creature.Name)
	c.creature.Attack *= 2
	c.CreatureModifier.Handle()
}

func main() {
	g := NewCreature("Goblin", 1, 1)
	fmt.Println(g.String())

	root := NewCreatureModifier(g)
	root.Add(NewDoubleAttackModifier(g))
	root.Handle()

	fmt.Println(g.String())
}
