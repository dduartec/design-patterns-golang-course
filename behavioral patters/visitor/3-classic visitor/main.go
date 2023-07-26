package main

import (
	"fmt"
	"strings"
)

// if another experssion is added, a visit method must be created
type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(e ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(e ExpressionVisitor) {
	e.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(e ExpressionVisitor) {
	e.VisitAdditionExpression(a)
}

// implements visitor
type ExpressionPrinter struct {
	s *strings.Builder
}

func (ep ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.s.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.s.WriteRune('(')
	e.left.Accept(ep)
	ep.s.WriteRune('+')
	e.right.Accept(ep)
	ep.s.WriteRune(')')
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{&strings.Builder{}}
}

func (ep *ExpressionPrinter) String() string {
	return ep.s.String()
}

// Add another visitor y much more easier
type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	l := ee.result
	e.right.Accept(ee)
	r := ee.result
	ee.result = l + r
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
	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g \n", ep, ee.result)
}
