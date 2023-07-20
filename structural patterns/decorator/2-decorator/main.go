package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float64
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float64) {
	c.Radius *= factor
}

type Square struct {
	Side float64
}

func (c *Square) Render() string {
	return fmt.Sprintf("Square with side %f", c.Side)
}

// DEcorator
type ColoredShape struct {
	Shape Shape
	Color string
}

func (cs *ColoredShape) Render() string {
	return fmt.Sprintf("%s has color %s", cs.Shape.Render(), cs.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float64
}

func (ts *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %d%% transparency", ts.Shape.Render(), int(ts.Transparency*100.0))
}

func main() {
	c := &Circle{2}
	c.Resize(2)
	fmt.Println(c.Render())

	// add functionality to the original
	rc := ColoredShape{c, "red"}
	fmt.Println(rc.Render())
	// Limitation: cant be resized
	// rc.Resize(2)

	// add another functionality
	tc := TransparentShape{&rc, 0.5}
	fmt.Println(tc.Render())

}
