#color

```
import (
	"image/color"
	"fmt"
	"github.com/xyproto/color"
)

func main() {

	hbs := color.RGBToHLS(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Print(hbs.Gethsb())

	rgb := hbs.HSB2RGB()
	fmt.Println(rgb)
}
```
