package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.width, im.height)
}

func (Image) At(x, y int) color.Color {
	v := uint8(x*x) + uint8(y*y)
	return color.RGBA{R: 255, G: v, B: v, A: 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
