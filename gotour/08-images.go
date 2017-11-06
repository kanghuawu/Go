package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
    return color.GrayModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 255, 255)
}

func (i Image) At(x, y int) color.Color {
    return color.Gray{uint8(x ^ y)}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}