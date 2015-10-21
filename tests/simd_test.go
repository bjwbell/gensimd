// +build amd64,gc

package tests

import (
	"math"
	"math/rand"
	"testing"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -debug -fn "addi8x16, subi8x16, addu8x16, subu8x16" -outfn "addi8x16s, subi8x16s, addu8x16s, subu8x16s" -f "$GOFILE" -o "simd_test_amd64.s"

func addi8x16s(x, y simd.I8x16) simd.I8x16
func subi8x16s(x, y simd.I8x16) simd.I8x16

func addu8x16s(x, y simd.U8x16) simd.U8x16
func subu8x16s(x, y simd.U8x16) simd.U8x16

func addi8x16(x, y simd.I8x16) simd.I8x16 { return simd.AddI8x16(x, y) }
func subi8x16(x, y simd.I8x16) simd.I8x16 { return simd.SubI8x16(x, y) }

func addu8x16(x, y simd.U8x16) simd.U8x16 { return simd.AddU8x16(x, y) }
func subu8x16(x, y simd.U8x16) simd.U8x16 { return simd.SubU8x16(x, y) }

func TestSimd(t *testing.T) {

	count := 0
	rand.Seed(42)

	for i := -63; i <= 63; i++ {

		a := int(0)
		if i < 0 {
			a = -1 << uint(-i)
		} else {
			a = 1<<uint(i) - 1
		}

		for j := -63; j <= 63; j++ {

			count++

			b := int(0)
			if j < 0 {
				b = -1 << uint(-j)
			} else {
				b = 1<<uint(j) - 1
			}

			var x simd.I8x16
			var y simd.I8x16
			var xu simd.U8x16
			var yu simd.U8x16

			for idx := 0; idx < 16; idx++ {
				abs_a := a

				if abs_a == math.MinInt64 || abs_a == 0 {
					abs_a = 64
				}
				if abs_a < 0 {
					abs_a = -abs_a
				}
				abs_b := b
				if abs_b == math.MinInt64 || abs_b == 0 {
					abs_b = 64
				}
				if abs_b < 0 {
					abs_b = -abs_b
				}

				x[idx] = int8(rand.Intn(abs_a))
				y[idx] = int8(rand.Intn(abs_b))

				xu[idx] = uint8(rand.Intn(abs_a))
				yu[idx] = uint8(rand.Intn(abs_b))
			}

			if addi8x16s(x, y) != addi8x16(x, y) {
				t.Errorf("addi8x16(%v, %v)", x, y)
				t.Error("x:", x)
				t.Error("y:", y)
				t.Error("s(x, y):", addi8x16s(x, y))
				t.Error(" (x, y):", addi8x16(x, y))
			}
			if subi8x16s(x, y) != subi8x16(x, y) {
				t.Errorf("subi8x16(%v, %v)", x, y)
				t.Error("x:", x)
				t.Error("y:", y)
				t.Error("s:", subi8x16s(x, y))
				t.Error(" :", subi8x16(x, y))
			}
			if addu8x16s(xu, yu) != addu8x16(xu, yu) {
				t.Errorf("addu8x16(%v, %v)", a, b)
			}
			if subu8x16s(xu, yu) != subu8x16(xu, yu) {
				t.Errorf("subu8x16(%v, %v)", xu, yu)
				t.Error("x:", xu)
				t.Error("y:", yu)
				t.Error("s:", subu8x16s(xu, yu))
				t.Error(" :", subu8x16(xu, yu))

			}
		}
	}

	t.Log("Test Count:", count)
}
