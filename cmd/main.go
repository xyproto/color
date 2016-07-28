package main

import (
	"fmt"
	hbscolor "github.com/xyproto/color"
	"image/color"
)

func main() {

	hbs := hbscolor.RGBToHLS(color.RGBA{R: 153, G: 153, B: 153})
	fmt.Print(hbs.Gethsb())

	rgb := hbs.HSB2RGB()
	fmt.Println(rgb)
}
