package color

import (

	"image/color"
	"math"
)

type HSB struct {
	h int //色相
	s, b float64// ,饱和度(纯度),l/b亮度
}

func (hsb HSB)Gethsb()(int,int,int){
	return hsb.h,int(hsb.s*100),int(hsb.b*100)
}

func RGBToHLS(rgb color.RGBA) (hsb HSB) {

	var r, g, b float64
	var rgb_min, rgb_max, delta float64

	r = float64(rgb.R) / 255.0
	g = float64(rgb.G) / 255.0
	b = float64(rgb.B) / 255.0

	rgb_max = math.Max(r, math.Max(g, b))
	rgb_min = math.Min(r, math.Min(g, b))
	delta = rgb_max - rgb_min

	if delta== 0 {
		hsb.h = 0
		hsb.s = 0
	} else if rgb_max == r {
		if g >= b {
			hsb.h = int(60 * (g - b) / delta)
		} else {
			hsb.h = int(60*(g-b)/delta + 360)
		}
	} else if rgb_max == g {
		hsb.h = int(60*(b-r)/delta + 120)
	} else {
		hsb.h = int(60*(r-g)/delta + 240)
	}

	hsb.b = rgb_max

	if hsb.b == 0 {
		hsb.s = 0
	} else if hsb.b <=0.5 {
		hsb.s = delta*2 / (2.0 *hsb.b)
	} else {
		hsb.s = delta *2 / (2.0 - 2.0*hsb.b)
	}
	return
}

func (hsl HSB) HSB2RGB() (color.RGBA) {
	var max_rgb_val = 255.0
	var sat = hsl.s*100
	var bright = hsl.b*100
	var hue = hsl.h
	var sat_f = float64(sat) / 100.0
	var bright_f = float64(bright) / 100.0
	var r, g, b int

	var colors [3]int

	if bright <= 0 {
		colors[0] = 0
		colors[1] = 0
		colors[2] = 0
	}
	if sat <= 0 {
		colors[0] = int(bright_f * max_rgb_val)
		colors[1] = int(bright_f * max_rgb_val)
		colors[2] = int(bright_f * max_rgb_val)
	} else {
		if hue >= 0 && hue < 120 {
			var hue_primary = 1.0 - (float64(hue) / 120.0)
			var hue_secondary = float64(hue) / 120.0
			var sat_primary = (1.0 - hue_primary) * (1.0 - sat_f)
			var sat_secondary = (1.0 - hue_secondary) * (1.0 - sat_f)
			var sat_tertiary = 1.0 - sat_f
			r = int((bright_f * max_rgb_val) * (hue_primary + sat_primary))
			g = int((bright_f * max_rgb_val) * (hue_secondary + sat_secondary))
			b = int((bright_f * max_rgb_val) * sat_tertiary)
		} else if hue >= 120 && hue < 240 {
			var hue_primary = 1.0 - ((float64(hue) - 120.0) / 120.0)
			var hue_secondary = (float64(hue) - 120.0) / 120.0
			var sat_primary = (1.0 - hue_primary) * (1.0 - sat_f)
			var sat_secondary = (1.0 - hue_secondary) * (1.0 - sat_f)
			var sat_tertiary = 1.0 - sat_f
			r = int((bright_f * max_rgb_val) * sat_tertiary)
			g = int((bright_f * max_rgb_val) * (hue_primary + sat_primary))
			b = int((bright_f * max_rgb_val) * (hue_secondary + sat_secondary))
		} else if hue >= 240 && hue <= 360 {
			var hue_primary = 1.0 - ((float64(hue) - 240.0) / 120.0)
			var hue_secondary = (float64(hue) - 240.0) / 120.0
			var sat_primary = (1.0 - hue_primary) * (1.0 - sat_f)
			var sat_secondary = (1.0 - hue_secondary) * (1.0 - sat_f)
			var sat_tertiary = 1.0 - sat_f
			r = int((bright_f * max_rgb_val) * (hue_secondary + sat_secondary))
			g = int((bright_f * max_rgb_val) * sat_tertiary)
			b = int((bright_f * max_rgb_val) * (hue_primary + sat_primary))
		}

		colors[0] = r
		colors[1] = g
		colors[2] = b

	}
	return color.RGBA{R:uint8(r), G:uint8(g), B:uint8(b)}
}
