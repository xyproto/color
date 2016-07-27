// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"fmt"

	"testing"
)

func delta(x, y uint8) uint8 {
	if x >= y {
		return x - y
	}
	return y - x
}

func TestSpecific(t *testing.T) {
	//r0, g0, b0 := uint8(79), uint8(24), uint8(23)
	//h, l, s := RGBToHLS(r0, g0, b0)
	//fmt.Print(h*360,l*100,s*100)

	var hls = HSB{h: 1, b: 71, s: 31}

	fmt.Print(hls.HSB2RGB())
	//if delta(r0, r1) > 1 || delta(g0, g1) > 1 || delta(b0, b1) > 1 {
	//	t.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
	//}
}
