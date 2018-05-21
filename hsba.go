package color

import (
	"image/color"
	"math"
)

// HSBA stands for hue, saturation, brightness and alpha
type HSBA struct {
	H    int     // hue, from 0 to 360
	S, B float64 // saturation, brightness, from 0 to 1
	A    float64 // alpha, from 0 to 1
}

// AsInts returns the hue, saturation, brightness and alpha values as four integers, from 0 to 360, 0 to 255, 0 to 255 and 0 to 255
func (h *HSBA) AsInts() (int, int, int, int) {
	return h.H, int(h.S * 255), int(h.B * 255), int(h.A * 255)
}

// NewFromFloats takes hue, saturation, brightness and alpha as four floats and returns a pointer to an HSBA struct
func NewFromFloats(h, s, b, a float64) *HSBA {
	return &HSBA{int(h * 360.0), s, b, 1.0}
}

// New takes a color.RGBA struct and returns a pointer to an HSBA struct
func New(rgb color.RGBA) *HSBA {
	var r, g, b float64
	var rgb_min, rgb_max, delta float64

	r = float64(rgb.R) / 255.0
	g = float64(rgb.G) / 255.0
	b = float64(rgb.B) / 255.0

	rgb_max = math.Max(r, math.Max(g, b))
	rgb_min = math.Min(r, math.Min(g, b))

	h := &HSBA{0, 0.0, rgb_max, float64(rgb.A) / 255.0}

	delta = rgb_max - rgb_min

	if rgb_max == 0 {
		h.S = 0
	} else {
		h.S = delta / rgb_max
	}

	if delta == 0 {
		h.H = 0
		h.S = 0
	} else if rgb_max == r {
		if g >= b {
			h.H = int(60 * (g - b) / delta)
		} else {
			h.H = int(60*(g-b)/delta + 360)
		}
	} else if rgb_max == g {
		h.H = int(60*(b-r)/delta + 120)
	} else {
		h.H = int(60*(r-g)/delta + 240)
	}

	return h
}

// Return the color as HSB, without Alpha
func (h *HSBA) WithoutAlpha() *HSB {
	return &HSB{h.H, h.S, h.B}
}

// Return the color as RGBA bytes, in a color.RGBA struct
func (h *HSBA) RGBA() color.RGBA {
	r, g, b := 0.0, 0.0, 0.0
	i := int((h.H / 60) % 6)
	var f float64 = float64(h.H)/60.0 - float64(i)
	p := h.B * (1 - h.S)
	q := h.B * (1.0 - f*h.S)
	t := h.B * (1 - (1-f)*h.S)

	switch i {
	case 0:
		r = h.B
		g = t
		b = p
	case 1:
		r = q
		g = h.B
		b = p
	case 2:
		r = p
		g = h.B
		b = t
	case 3:
		r = p
		g = q
		b = h.B
	case 4:
		r = t
		g = p
		b = h.B
	case 5:
		r = h.B
		g = p
		b = q
	}
	return color.RGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: uint8(h.A * 255)}
}

// Return the color as RGB bytes, in a color.RGBA struct
func (h *HSBA) RGB() color.RGBA {
	rgba := h.RGBA()
	rgba.A = 255
	return rgba
}
