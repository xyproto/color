# color

## Example usage:

```
package main

import (
	"fmt"
	hbscolor "github.com/xyproto/color"
	"image/color"
)

func main() {

	hbs := hbscolor.RGBToHLS(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Print(hbs.Gethsb())

	rgb := hbs.HSB2RGB()
	fmt.Println(rgb)

}
```
