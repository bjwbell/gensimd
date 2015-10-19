// +build amd64,gc

package tests

import (
	"math"
	"testing"
)

//go:generate gensimd -fn "ift0, ift1, ift2, ift3, ift4, ift5, ift6, ift7, ift8, ift9" -outfn "ift0s, ift1s, ift2s, ift3s, ift4s, ift5s, ift6s, ift7s, ift8s, ift9s" -f "$GOFILE" -o "if_test_amd64.s"

func ift0s(uint8) uint8
func ift1s(uint16) uint16
func ift2s(uint32) uint32
func ift3s(uint64) uint64
func ift4s(int8) int8
func ift5s(int16) int16
func ift6s(int32) int32
func ift7s(int64) int64
func ift8s(float32) float32
func ift9s(float64) float64

func ift0(x uint8) uint8 {
	if x < 2 {
		return x
	} else {
		return x * x
	}
}
func ift1(x uint16) uint16 {
	if x > 128 {
		return ^x
	} else {
		return x
	}
}
func ift2(x uint32) uint32 {
	if x < 1024 {
		return x & 509
	} else {
		return x & 511
	}
}
func ift3(x uint64) uint64 {
	if x*x < 2046 {
		return x
	} else {
		return 2*x - x
	}
}
func ift4(x int8) int8 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func ift5(x int16) int16 {
	if x < -255 {
		return -255 * x
	} else {
		return 255 * x
	}
}
func ift6(x int32) int32 {
	if x == 1 {
		return 0
	} else {
		return 1
	}
}
func ift7(x int64) int64 {
	if x < -1 {
		return -x
	} else {
		return 10 - x
	}
}
func ift8(x float32) float32 {
	if x < 1 {
		return -x
	} else {
		return 10 - x
	}
}
func ift9(x float64) float64 {
	if x*x < 2046 {
		return x
	} else {
		return 3*x - x
	}
}

func TestIf(t *testing.T) {

	count := 0

	for j := -63; j <= 63; j++ {

		count++

		x := int64(0)
		if j < 0 {
			x = -1 << uint(-j)
		} else {
			x = 1<<uint(j) - 1
		}

		if ift0s(uint8(x)) != ift0(uint8(x)) {
			t.Errorf("ift0s (%v) != ift0 (%v)", ift0s(uint8(x)), ift0(uint8(x)))
		}
		if ift1s(uint16(x)) != ift1(uint16(x)) {
			t.Errorf("ift1s (%v) != ift1 (%v)", ift1s(uint16(x)), ift1(uint16(x)))
		}
		if ift2s(uint32(x)) != ift2(uint32(x)) {
			t.Errorf("ift2s (%v) != ift2 (%v)", ift2s(uint32(x)), ift2(uint32(x)))
		}
		if ift3s(uint64(x)) != ift3(uint64(x)) {
			t.Errorf("ift3s (%v) != ift3 (%v)", ift3s(uint64(x)), ift3(uint64(x)))
		}

		if ift4s(int8(x)) != ift4(int8(x)) {
			t.Errorf("ift4s (%v) != ift4 (%v)", ift4s(int8(x)), ift4(int8(x)))
		}
		if ift5s(int16(x)) != ift5(int16(x)) {
			t.Errorf("ift5s (%v) != ift5 (%v)", ift5s(int16(x)), ift5(int16(x)))
		}
		if ift6s(int32(x)) != ift6(int32(x)) {
			t.Errorf("ift6s (%v) != ift6 (%v)", ift6s(int32(x)), ift6(int32(x)))
		}
		if ift7s(x) != ift7(x) {
			t.Errorf("ift7s (%v) != ift7 (%v)", ift7s(x), ift7(x))
		}

		if ift8s(float32(x)) != ift8(float32(x)) {
			sbits := math.Float32bits(ift8s(float32(x)))
			tbits := math.Float32bits(ift8(float32(x)))
			t.Errorf("ift8s (x=%v) (%v=%08x) != ift8 (%v=%08x)", float32(x), ift8s(float32(x)), sbits, ift8(float32(x)), tbits)
		}
		if ift9s(float64(x)) != ift9(float64(x)) {
			t.Errorf("ift9s (%v) != ift9 (%v)", ift9s(float64(x)), ift9(float64(x)))
		}
	}

	t.Log("Test Count:", count)
}
