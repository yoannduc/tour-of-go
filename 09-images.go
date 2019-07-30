package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Define Image struct
type Image struct {
	Width  int
	Height int
}

// Implement Bounds method as per the Image interface
func (m *Image) Bounds() image.Rectangle {
	// Return an image rectangle with struct width & height
	return image.Rect(0, 0, m.Width, m.Height)
}

// Implement ColorModel function as per the Image interface
func (m *Image) ColorModel() color.Model {
	// Return a RGBA model
	return color.RGBAModel
}

// Implement the At function as per the Image interface
func (m *Image) At(x, y int) color.Color {
	// Tweak the image with ^ operator
	c := uint8(x ^ y)
	// Return a RGBA color with tweaked RG
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := &Image{256, 256}
	pic.ShowImage(m)
}
