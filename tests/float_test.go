// +build amd64,gc

package tests

import (
	"math/rand"
	"testing"
)

//go:generate gensimd -fn "addf32, subf32, negf32, mulf32, addf64, subf64, negf64, mulf64" -outfn "addf32s, subf32s, negf32s, mulf32s, addf64s, subf64s, negf64s, mulf64s" -f "$GOFILE" -o "float_test.s"

func addf32s(x, y float32) float32
func subf32s(x, y float32) float32
func negf32s(x float32) float32
func mulf32s(x, y float32) float32

//func divf32s(x, y float32) float32

func addf64s(x, y float64) float64
func subf64s(x, y float64) float64
func negf64s(x float64) float64
func mulf64s(x, y float64) float64

//func divf64s(x, y float64) float64

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

/*func divf32(x, y float32) float32 {
	return x / y
}*/
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

/*func divf64(x, y float64) float64 {
	return x / y
}*/

func TestFloatOps(t *testing.T) {

	for i := 0; i <= 128*128; i++ {
		var y float64
		if i&1 == 1 {
			y = rand.Float64()
		} else {
			y = rand.ExpFloat64()
		}
		if negf32s(float32(y)) != negf32(float32(y)) {
			t.Errorf("negf32s(%v)", float32(y))
		}
		if negf64s(float64(y)) != negf64(float64(y)) {
			t.Errorf("negf64s(%v)", float64(y))
		}

		for j := 0; j <= 256; j++ {
			var x float64
			if j&1 == 1 {
				x = rand.Float64()
			} else {
				x = rand.ExpFloat64()
			}

			if addf32s(float32(x), float32(y)) != addf32(float32(x), float32(y)) {
				t.Errorf("addf32s(%v, %v)", float32(x), float32(y))
			}
			if subf32s(float32(x), float32(y)) != subf32(float32(x), float32(y)) {
				t.Errorf("subf32s(%v, %v)", float32(x), float32(y))

			}
			if mulf32s(float32(x), float32(y)) != mulf32(float32(x), float32(y)) {
				t.Errorf("mulf32s(%v, %v)", float32(x), float32(y))
			}

			if addf64s(float64(x), float64(y)) != addf64(float64(x), float64(y)) {
				t.Errorf("addf64s(%v, %v)", float64(x), float64(y))
			}
			if subf64s(float64(x), float64(y)) != subf64(float64(x), float64(y)) {
				t.Errorf("subf64s(%v, %v)", float64(x), float64(y))

			}
		}
	}
}
