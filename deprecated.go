package color

import (
	"image/color"
)

// HSB stands for hue, saturation and brightness
type HSB struct {
	H    int     // 色相
	S, B float64 // ,饱和度(纯度),l/b亮度
}

// RGBToHLS is deprecated, use New instead
func RGBToHLS(rgb color.RGBA) HSB {
	hsb := New(rgb).WithoutAlpha()
	return *hsb
}

// Gethsb is deprecated, use hsba AsInts() instead
func (h HSB) Gethsb() (int, int, int) {
	return h.AsInts()
}

// HSB2RGB is deprecated, use hsba RGBA() instead
func (hsb HSB) HSB2RGB() color.RGBA {
	h := &HSBA{hsb.H, hsb.S, hsb.B, 1.0}
	return h.RGBA()
}

// HSB2RGB is deprecated, use hsba RGBA() instead
func (hsb HSB) RGB() color.RGBA {
	return hsb.HSB2RGB()
}

// AsInts is deprecated, use hsba AsInts instead
func (h *HSB) AsInts() (int, int, int) {
	return h.H, int(h.S * 255), int(h.B * 255)
}

func NewHSB(c color.RGBA) HSB {
	hsba := New(c)
	var hsb HSB
	hsb.H = hsba.H
	hsb.S = hsba.S
	hsb.B = hsba.B
	return hsb
}
