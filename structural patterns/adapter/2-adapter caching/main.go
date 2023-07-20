package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

type Line struct {
	x1, y1, x2, y2 int
}

type VectorImage struct {
	Line []Line
}

// Given interface
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		Line: []Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// interfce we have
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	return fmt.Sprintf("interface imp %d", len(owner.GetPoints()))
}

// Adapter
type vectorToRasterAdapter struct {
	points []Point
}

var pointCache = map[[16]byte][]Point{}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v *vectorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}
	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		for _, p := range pts {
			v.points = append(v.points, p)
		}
		fmt.Println("used cache")
		return
	}
	v.addLine(line)
	pointCache[h] = v.points
	fmt.Println("added cache")
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	v.points = append(v.points, Point{line.x1, line.y1}, Point{line.x2, line.y2})
}

// Adapt
func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Line {
		adapter.addLineCached(line)
	}

	return &adapter // As RasterImage
}

func main() {
	// origin interface
	rc := NewRectangle(6, 4)
	// adapto to destination interface
	a := VectorToRaster(rc)
	_ = VectorToRaster(rc) // another adapter
	// use adapted
	fmt.Println(DrawPoints(a))
}
