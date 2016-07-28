package color

import (
	"image/color"
	"math"
)

type HSB struct {
	H    int     //色相
	S, B float64 // ,饱和度(纯度),l/b亮度
}

func (hsb HSB) Gethsb() (int, int, int) {
	return hsb.H, int(hsb.S * 100), int(hsb.B * 100)
}

func RGBToHLS(rgb color.RGBA) (hsb HSB) {

	var r, g, b float64
	var rgb_min, rgb_max, delta float64

	r = float64(rgb.R) / 255.0
	g = float64(rgb.G) / 255.0
	b = float64(rgb.B) / 255.0

	rgb_max = math.Max(r, math.Max(g, b))
	rgb_min = math.Min(r, math.Min(g, b))
	hsb.B = rgb_max
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

	return
}

func (hsl HSB) HSB2RGB() color.RGBA {
	r, g, b := 0.0, 0.0, 0.0
	i := int((hsl.H / 60) % 6)
	var f float64 = float64(hsl.H)/60.0 - float64(i)
	p := hsl.B * (1 - hsl.S)
	q := hsl.B * (1.0 - f*hsl.S)
	t := hsl.B * (1 - (1-f)*hsl.S)

	switch i {
	case 0:
		r = hsl.B
		g = t
		b = p
	case 1:
		r = q
		g = hsl.B
		b = p
	case 2:
		r = p
		g = hsl.B
		b = t
	case 3:
		r = p
		g = q
		b = hsl.B
	case 4:
		r = t
		g = p
		b = hsl.B
	case 5:
		r = hsl.B
		g = p
		b = q
	}
	return color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255)}
}
