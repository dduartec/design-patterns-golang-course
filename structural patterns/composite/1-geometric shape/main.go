package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (o *GraphicObject) AddChild(child GraphicObject) {
	o.Children = append(o.Children, child)
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

// Use the same logic with scalars and lists
func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	// if has children then use recursion
	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func main() {
	drawing := GraphicObject{"Drawing", "", nil}
	drawing.AddChild(*NewCircle("red"))
	drawing.AddChild(*NewSquare("blue"))

	group := GraphicObject{"Group", "", nil}
	group.AddChild(*NewCircle("green"))
	group.AddChild(*NewSquare("green"))

	drawing.AddChild(group)

	fmt.Println(drawing.String())
}
