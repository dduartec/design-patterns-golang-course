package main

import (
	"fmt"
	"sync"
)

// CoR, Mediator, Observer, CommandQuerySeparation(CQS)

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	//CoR: run queries in order
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		fmt.Println("Aplying query over", q.CreatureName, q.)
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game // mediator
	Name            string
	attack, defense int
}

func NewCreature(game *Game, name string, attack, defense int) *Creature {
	return &Creature{game, name, attack, defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s(%d/%d)", c.Name, c.Attack(), c.Defense())
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (cm *CreatureModifier) Handle(q *Query) {

}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

func (cm *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == cm.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}
func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func main() {
	game := &Game{sync.Map{}}
	g1 := NewCreature(game, "goblin", 1, 1)
	g := NewCreature(game, "Stromg goblin", 2, 2)
	{
		m := NewDoubleAttackModifier(game, g)
		fmt.Println(g)
		_ = NewDoubleAttackModifier(game, g1)
		fmt.Println(g1)
		m.Close()
	}
	fmt.Println(g)
	fmt.Println(g1)
}
