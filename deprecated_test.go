package color

import (
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
