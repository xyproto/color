package color

import (
	"fmt"
	"image/color"
	"testing"
)

func TestHSB(t *testing.T) {
	hsb := HSB{H: 1, B: 71, S: 31}
	rgb := hsb.HSB2RGB()
	if rgb.R != byte(185) {
		t.Fatalf("%d != %d", rgb.R, 185)
	}
	if rgb.G != byte(221) {
		t.Fatalf("%d != %d", rgb.G, 221)
	}
	if rgb.B != byte(82) {
		t.Fatalf("%d != %d", rgb.B, 82)
	}
}

func ExampleHSB() {
	hsb := NewHSB(color.RGBA{R: 79, G: 24, B: 23})
	fmt.Println(hsb.AsInts())
	// Output: 1 180 79
}
