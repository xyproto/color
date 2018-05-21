# color

* HSB stands for Hue, Saturation and Brightness
* RGB stands for Red, Green and Blue

## Example usage:

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

	rgb := hsb.RGB()
	fmt.Println(rgb)
}
```
