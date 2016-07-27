package main

import (
	"image/color"
	"fmt"
	hbscolor "github.com/mengxiaozhu/color"
)

func main() {

	hbs :=hbscolor.RGBToHLS(color.RGBA{R:255, G: 255, B: 23})
	fmt.Print(hbs.Gethsb())

	rgb:=hbs.HSB2RGB()
	fmt.Println(rgb)
}
