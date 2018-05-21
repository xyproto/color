package main

import (
	"fmt"
	hbscolor "github.com/xyproto/color"
	"image/color"
)

func main() {
	hsb := hbscolor.New(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Println(hsb.AsInts())
	// Outputs: 1 70 30

	rgb := hsb.RGB()
	fmt.Println(rgb)
	// Outputs: {79 23 23 0}
}
