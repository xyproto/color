package color

import (
	"image/color"
)

// Gethsb is deprecated, use hsb.AsInts() instead
func (hsb HSB) Gethsb() (int, int, int) {
	return hsb.AsInts()
}

// RGBToHLS is deprecated, use New instead
func RGBToHLS(rgb color.RGBA) HSB {
	hsb := New(rgb)
	return *hsb
}

// HSB2RGB is deprecated, use hsb.RGB() instead
func (hsb HSB) HSB2RGB() color.RGBA {
	return hsb.RGB()
}
