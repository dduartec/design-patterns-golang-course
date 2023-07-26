package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinnigPlayer() int
}

// template method
func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins. \n", g.WinnigPlayer())
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{1, 10, 0}
}

func (c *chess) Start() {
	fmt.Println("Starting new game of chess")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c chess) WinnigPlayer() int {
	return c.currentPlayer
}

func main() {
	c := NewGameOfChess()
	PlayGame(c)
}
