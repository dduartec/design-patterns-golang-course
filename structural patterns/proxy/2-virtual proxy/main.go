package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image", filename)
	return &Bitmap{filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Finished drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (b *LazyBitmap) Draw() {
	if b.bitmap == nil {
		b.bitmap = NewBitmap(b.filename)
	}
	b.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	// Load the image even if we don't use it
	// bmp := NewBitmap("fdafdas")
	// Only loads the image when needed 👍
	bmp := NewLazyBitmap("fdasfdsa")
	DrawImage(bmp)
}
