package main

import "fmt"

//OCP
// Open for extendion, closed for modification
// Specification pattern

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Does not follow OCP
type Filter struct{}

// If we want to add more filters, we would need to add more methods.
// This is not convinient because the logic would be the same but for another attribute
func (f *Filter) FilterByColor(products []Product, color Color) []Product {
	result := make([]Product, 0)
	for _, v := range products {
		if v.color == color {
			result = append(result, v)
		}
	}
	return result
}

type Specification interface {
	isSatisfied(p Product) bool
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) isSatisfied(p Product) bool {
	return a.first.isSatisfied(p) && a.second.isSatisfied(p)
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) isSatisfied(p Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) isSatisfied(p Product) bool {
	return s.size == p.size
}

type BetterFilter struct{}

func (f BetterFilter) Filter(products []Product, spec Specification) []Product {
	result := make([]Product, 0)
	for _, v := range products {
		if spec.isSatisfied(v) {
			result = append(result, v)
		}
	}
	return result
}

var (
	greenSpec  ColorSpecification = ColorSpecification{green}
	redSpec    ColorSpecification = ColorSpecification{red}
	blueSpec   ColorSpecification = ColorSpecification{blue}
	smallSpec  SizeSpecification  = SizeSpecification{small}
	mediumSpec SizeSpecification  = SizeSpecification{medium}
	largeSpec  SizeSpecification  = SizeSpecification{large}
)

func main() {
	apple := Product{
		name:  "Apple",
		color: red,
		size:  small,
	}
	tree := Product{
		name:  "Tree",
		color: green,
		size:  medium,
	}
	house := Product{
		name:  "House",
		color: blue,
		size:  large,
	}
	car := Product{
		name:  "Car",
		color: blue,
		size:  large,
	}
	products := []Product{apple, tree, house, car}
	f := Filter{}
	redProducts := f.FilterByColor(products, red)
	fmt.Println(redProducts)

	bf := BetterFilter{}
	greenProds := bf.Filter(products, greenSpec)
	fmt.Println(greenProds)
	blueLargeProds := bf.Filter(products, AndSpecification{blueSpec, largeSpec})
	fmt.Println(blueLargeProds)

}
