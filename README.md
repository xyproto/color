# color

A package for converting from RGB to HSB and back.

* HSB stands for Hue, Saturation and Brightness
* RGB stands for Red, Green and Blue

## Example usage

```go
package main

import (
	"fmt"
	"image/color"
	hbscolor "github.com/xyproto/color"
)

func main() {
	hsb := hbscolor.New(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Println(hsb.AsInts())
	// Outputs: 1 70 30

	rgb := hsb.RGB()
	fmt.Println(rgb)
	// Outputs: {79 23 23 0}
}
```

# General info

* License: MIT
