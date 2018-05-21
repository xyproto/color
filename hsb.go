package color

import (
	"image/color"
	"math"
)

// HSB stands for hue, saturation and brightness
type HSB struct {
	H    int     // 色相
	S, B float64 // ,饱和度(纯度),l/b亮度
}

// AsInts returns the hue, saturation and brightness values as three integers, from 0 to 100
func (hsb *HSB) AsInts() (int, int, int) {
	return hsb.H, int(hsb.S * 100), int(hsb.B * 100)
}

// New takes a color.RGBA struct and returns a pointer to an HSB struct
func New(rgb color.RGBA) *HSB {
	var r, g, b float64
	var rgb_min, rgb_max, delta float64

	r = float64(rgb.R) / 255.0
	g = float64(rgb.G) / 255.0
	b = float64(rgb.B) / 255.0

	rgb_max = math.Max(r, math.Max(g, b))
	rgb_min = math.Min(r, math.Min(g, b))

	hsb := &HSB{0, 0.0, rgb_max}

	delta = rgb_max - rgb_min

	if rgb_max == 0 {
		hsb.S = 0
	} else {
		hsb.S = delta / rgb_max
	}

	if delta == 0 {
		hsb.H = 0
		hsb.S = 0
	} else if rgb_max == r {
		if g >= b {
			hsb.H = int(60 * (g - b) / delta)
		} else {
			hsb.H = int(60*(g-b)/delta + 360)
		}
	} else if rgb_max == g {
		hsb.H = int(60*(b-r)/delta + 120)
	} else {
		hsb.H = int(60*(r-g)/delta + 240)
	}

	return hsb
}

// Return the color as RGB bytes, in a color.RGBA struct
func (hsb *HSB) RGB() color.RGBA {
	r, g, b := 0.0, 0.0, 0.0
	i := int((hsb.H / 60) % 6)
	var f float64 = float64(hsb.H)/60.0 - float64(i)
	p := hsb.B * (1 - hsb.S)
	q := hsb.B * (1.0 - f*hsb.S)
	t := hsb.B * (1 - (1-f)*hsb.S)

	switch i {
	case 0:
		r = hsb.B
		g = t
		b = p
	case 1:
		r = q
		g = hsb.B
		b = p
	case 2:
		r = p
		g = hsb.B
		b = t
	case 3:
		r = p
		g = q
		b = hsb.B
	case 4:
		r = t
		g = p
		b = hsb.B
	case 5:
		r = hsb.B
		g = p
		b = q
	}
	return color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255)}
}
