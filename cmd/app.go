package main

import (
	"image/color"
	"fmt"
	c "github.com/mengxiaozhu/color"
)

func main() {

	hbs :=c.RGBToHLS(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Print(hbs)
}
