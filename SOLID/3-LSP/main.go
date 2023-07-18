package main

// Liskov Substitution Principle
// If an API works correctly with your base class
// it should also runs correctly with the derive class

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

// We broke the assumptions about the parent type
// beaking LSP
func (r *Square) SetWidth(width int) {
	r.width = width
	r.height = width
}

func (r *Square) SetHeight(height int) {
	r.height = height
	r.width = height
}

func NewSquare(size int) *Square {
	return &Square{
		Rectangle{
			width:  size,
			height: size,
		},
	}
}

// Does not follow LSP
func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	area := sized.GetHeight() * sized.GetWidth()
	fmt.Println("Expected", expectedArea, "Actual", area)
}

func main() {
	r := &Rectangle{
		width:  2,
		height: 2,
	}
	UseIt(r)
	s := NewSquare(2)
	UseIt(s)
}
