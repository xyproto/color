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
	var hls = HSB{H: 1, B: 71, S: 31}
	fmt.Print(hls.HSB2RGB())
}
