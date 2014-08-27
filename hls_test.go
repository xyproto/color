// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"testing"
)

func delta(x, y uint8) uint8 {
	if x >= y {
		return x - y
	}
	return y - x
}

func TestSpecific(t *testing.T) {
	r0, g0, b0 := uint8(7), uint8(0), uint8(0)
	h, l, s := RGBToHLS(r0, g0, b0)
	r1, g1, b1 := HLSToRGB(h, l, s)
	if delta(r0, r1) > 1 || delta(g0, g1) > 1 || delta(b0, b1) > 1 {
		t.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
	}
}

// Test that a subset of RGB space can be converted to HLS and back to within
// 1/256 tolerance.
//func TestRoundtrip(t *testing.T) {
//	for r := 0; r < 255; r += 7 {
//		for g := 0; g < 255; g += 5 {
//			for b := 0; b < 255; b += 3 {
//				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
//				h, l, s := RGBToHLS(r0, g0, b0)
//				r1, g1, b1 := HLSToRGB(h, l, s)
//				if delta(r0, r1) > 1 || delta(g0, g1) > 1 || delta(b0, b1) > 1 {
//					t.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
//				}
//			}
//		}
//	}
//}
