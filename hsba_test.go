package color

import (
	"fmt"
)

func ExampleHSBA() {
	h := &HSBA{H: 1, B: 71, S: 31, A: 1.0}
	rgba := h.RGBA()
	fmt.Println(rgba)
	// Output: {185 221 82 255}
}
