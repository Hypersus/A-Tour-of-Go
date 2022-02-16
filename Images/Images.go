package main

import "golang.org/x/tour/pic"
import (
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,100,100)
}

func (i Image) At(x,y int) color.Color {
	c:=color.RGBA{uint8(x),uint8(y),255,255}
	return c
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

