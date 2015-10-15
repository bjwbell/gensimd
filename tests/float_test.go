// +build amd64,gc

package tests

import (
	"math/rand"
	"testing"
)

//go:generate gensimd -fn "ptrt0, ptrt1, addf32, subf32, negf32, mulf32, divf32, addf64, subf64, negf64, mulf64, divf64" -outfn "ptrt0s, ptrt1s, addf32s, subf32s, negf32s, mulf32s, divf32s, addf64s, subf64s, negf64s, mulf64s, divf64s" -f "$GOFILE" -o "float_test.s"

func ptrt0s(*float32) float32
func ptrt1s(*float64) float64

func addf32s(x, y float32) float32
func subf32s(x, y float32) float32
func negf32s(x float32) float32
func mulf32s(x, y float32) float32
func divf32s(x, y float32) float32

func addf64s(x, y float64) float64
func subf64s(x, y float64) float64
func negf64s(x float64) float64
func mulf64s(x, y float64) float64
func divf64s(x, y float64) float64

func addf32(x, y float32) float32 {
	return x + y
}

func subf32(x, y float32) float32 {
	return x - y
}
func negf32(x float32) float32 {
	return -x
}
func mulf32(x, y float32) float32 {
	return x * y
}

func divf32(x, y float32) float32 {
	return x / y
}
func addf64(x, y float64) float64 {
	return x + y
}
func subf64(x, y float64) float64 {
	return x - y
}
func negf64(x float64) float64 {
	return -x
}
func mulf64(x, y float64) float64 {
	return x * y
}

func divf64(x, y float64) float64 {
	return x / y
}

func ptrt0(x *float32) float32 {
	return 2.0 * *x
}

func ptrt1(x *float64) float64 {
	return 2.0**x + *x
}

func TestFloatOps(t *testing.T) {

	for i := 0; i <= 128*128; i++ {
		var y float64
		if i&1 == 1 {
			y = rand.Float64()
		} else {
			y = rand.ExpFloat64()
		}
		yf32 := float32(y)

		if ptrt0(&yf32) != ptrt0s(&yf32) {
			t.Errorf("ptrt0s(%v)", yf32)
		}
		if ptrt1(&y) != ptrt1s(&y) {
			t.Errorf("ptrt1s(%v)", y)
		}

		if negf32s(yf32) != negf32(yf32) {
			t.Errorf("negf32s(%v)", yf32)
		}
		if negf64s(y) != negf64(y) {
			t.Errorf("negf64s(%v)", y)
		}

		for j := 0; j <= 256; j++ {
			var x float64
			if j&1 == 1 {
				x = rand.Float64()
			} else {
				x = rand.ExpFloat64()
			}
			xf32 := float32(x)

			if addf32s(xf32, yf32) != addf32(xf32, yf32) {
				t.Errorf("addf32s(%v, %v)", xf32, yf32)
			}
			if subf32s(xf32, yf32) != subf32(xf32, yf32) {
				t.Errorf("subf32s(%v, %v)", xf32, yf32)

			}
			if mulf32s(xf32, yf32) != mulf32(xf32, yf32) {
				t.Errorf("mulf32s(%v, %v)", xf32, yf32)
			}
			if divf32s(xf32, yf32) != divf32(xf32, yf32) {
				t.Errorf("divf32s(%v, %v) %v != divf32 (%v)", xf32, yf32, divf32s(xf32, yf32), divf32(xf32, yf32))
			}

			if addf64s(x, y) != addf64(x, y) {
				t.Errorf("addf64s(%v, %v)", x, y)
			}
			if subf64s(x, y) != subf64(x, y) {
				t.Errorf("subf64s(%v, %v)", x, y)

			}
			if mulf64s(x, y) != mulf64(x, y) {
				t.Errorf("mulf64s(%v, %v)", x, y)
			}
			if divf64s(x, y) != divf64(x, y) {
				t.Errorf("divf64s(%v, %v) %v != divf64 (%v)", x, y, divf64s(x, y), divf64(x, y))
			}
		}
	}
}
