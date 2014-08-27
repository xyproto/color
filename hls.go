package color

import (
	"fmt"
	"image/color"
	"math"
)

func RGBToHLS(r0, g0, b0 uint8) (h float64, l float64, s float64) {

	r, g, b := float64(r0)/255.0, float64(g0)/255.0, float64(b0)/255.0

	fmt.Println("rgb", r, g, b) // 0.02 0 0

	maxc := math.Max(math.Max(r, g), b)
	minc := math.Min(math.Min(r, g), b)
	l = (minc + maxc) / 2.0
	if minc == maxc {
		return 0, l, 0
	}
	span := (maxc - minc)
	if l <= 0.5 {
		s = span / (maxc + minc)
	} else {
		s = span / (2.0 - maxc - minc)
	}
	rc := (maxc - r) / span
	gc := (maxc - g) / span
	bc := (maxc - b) / span
	if r == maxc {
		h = bc - gc
	} else if g == maxc {
		h = 2.0 + rc - bc
	} else {
		h = 4.0 + gc - rc
	}
	h = math.Mod((h / 6.0), 1.0)

	fmt.Println("hls", h, l, s)

	return h, l, s
}

func _v(m1, m2, hue float64) float64 {
	one_sixth := 1.0 / 6.0
	two_thirds := 2.0 / 3.0
	hue = math.Mod(hue, 1.0)
	if hue < one_sixth {
		return m1 + (m2-m1)*hue*6.0
	}
	if hue < 0.5 {
		return m2
	}
	if hue < two_thirds {
		return m1 + (m2-m1)*(two_thirds-hue)*6.0
	}
	return m1
}

func HLSToRGB(h, l, s float64) (uint8, uint8, uint8) {
	one_third := 1.0 / 3.0
	if s == 0 {
		lr := uint8(l * 255.0)
		return lr, lr, lr
	}
	var m2 float64
	if l <= 0.5 {
		m2 = l * (1.0 + s)
	} else {
		m2 = l + s - (l * s)
	}
	m1 := 2.0*l - m2
	r, g, b := _v(m1, m2, h+one_third), _v(m1, m2, h), _v(m1, m2, h-one_third)
	return uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)
}

type HLS struct {
	h, l, s float64
}

func (c HLS) RGBA() (uint32, uint32, uint32, uint32) {
	r, g, b := HLSToRGB(c.h, c.l, c.s)
	return uint32(r) * 0x101, uint32(g) * 0x101, uint32(b) * 0x101, 0xffff
}

// The Model for "Hue Lightness Saturation Alpha" colors.
var HLSModel color.Model = color.ModelFunc(hlsModel)

func hlsModel(c color.Color) color.Color {
	if _, ok := c.(HLS); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	h, l, s := RGBToHLS(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	return HLS{h, l, s}
}

// Mix two colors, using the HLS color system
func Mix(c1, c2 color.Color) color.Color {
	pi := 3.1415 // Using math.Pi gives a completely different result, for some reason

	var hlsa1, hlsa2 HLS

	if _, ok := c1.(HLS); ok {
		// c1 is already HLSA
		hlsa1 = c1.(HLS)
	} else {
		hlsa1 = hlsModel(c1).(HLS)
	}

	if _, ok := c2.(HLS); ok {
		// c2 is already HLSA
		hlsa2 = c2.(HLS)
	} else {
		hlsa1 = hlsModel(c2).(HLS)
	}

	h := 0.0
	s := (hlsa1.s + hlsa2.s) * 0.5
	l := (hlsa1.l + hlsa2.l) * 0.5
	x := math.Cos(2.0*pi*hlsa1.h) + math.Cos(2.0*pi*hlsa2.h)
	y := math.Sin(2.0*pi*hlsa1.h) + math.Sin(2.0*pi*hlsa2.h)

	if (x == 0) && (y == 0) {
		s = 0
	} else {
		h = math.Atan2(y, x) / (2.0 * pi)
	}

	return HLS{h, l, s}
}
