package main

import (
	"fmt"
	"strings"
)

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

// The visitor is external from the hierarchy: SRP not violated
func Print(e Expression, s *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		s.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		s.WriteRune('(')
		Print(ae.left, s)
		s.WriteRune('+')
		Print(ae.right, s)
		s.WriteRune(')')
	}
	// If a new expression is defined, more code must be writen
	// violates OCP
}

func main() {
	// 1+(2+3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	s := strings.Builder{}
	Print(e, &s)
	fmt.Println(s.String())
}
