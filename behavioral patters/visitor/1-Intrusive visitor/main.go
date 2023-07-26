package main

import (
	"fmt"
	"strings"
)

// the builder is the vistor, gets passed down over the hierarchy
// violates SRP
type Expression interface {
	Print(s *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(s *strings.Builder) {
	s.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(s *strings.Builder) {
	s.WriteRune('(')
	a.left.Print(s)
	s.WriteRune('+')
	a.right.Print(s)
	s.WriteRune(')')
}

func main() {
	// 1+(2+3)
	e := AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	s := strings.Builder{}
	e.Print(&s)
	fmt.Println(s.String())
}
