package main

import "fmt"

// Circle, squeare
// Raster, vector
// RasterCircle, RasterSquare, VectorCircle, VectorSquare ....

// Better

// Bridge interface
type Renderer interface {
	RenderCircle(radius float64)
}

type VectorRenderer struct {
	//
}

func (vr *VectorRenderer) RenderCircle(radius float64) {
	fmt.Println("Vector renderer circle of radius", radius)
}

type RasterRenderer struct {
	//
}

func (rr *RasterRenderer) RenderCircle(radius float64) {
	fmt.Println("Raster renderer circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius   float64
}

func (c Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{renderer, radius}
}

func main() {
	raster := RasterRenderer{}
	vec := VectorRenderer{}

	circle := NewCircle(&raster, 4)
	circle.Draw()

	circle.renderer = &vec
	circle.Draw()

}
