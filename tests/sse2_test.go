// +build amd64,gc

package tests

import (
	"math"
	"math/rand"
	"testing"

	"github.com/bjwbell/gensimd/simd"
	"github.com/bjwbell/gensimd/simd/sse2"
)

//go:generate gensimd -debug -fn "addpd_go" -outfn "addpd" -f "$GOFILE" -o "sse2_test_amd64.s"

func addpd(x, y simd.M128d) simd.M128d
func addpd_go(x, y simd.M128d) simd.M128d { return sse2.AddPd(x, y) }

func TestSSE2(t *testing.T) {
	// bail if sse2 is not supported
	if !simd.SSE2() {
		t.Log("Skipped - SSE2 not available")
		return
	}

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

			var xF64x2 simd.F64x2
			var yF64x2 simd.F64x2
			var xM128d simd.M128d
			var yM128d simd.M128d

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
				if idx < 2 {
					xF64x2[idx] = rand.ExpFloat64()
					yF64x2[idx] = rand.ExpFloat64()

					xM128d[idx] = xF64x2[idx]
					yM128d[idx] = yF64x2[idx]
				}

			}

			if addpd(xM128d, yM128d) != simd.M128d(simd.AddF64x2(xF64x2, yF64x2)) {
				t.Errorf("addm128d(%v, %v)", xF64x2, yF64x2)
				t.Error("x:", xF64x2)
				t.Error("y:", yF64x2)
				t.Error("s:", addpd(xM128d, yM128d))
				t.Error(" :", simd.AddF64x2(xF64x2, yF64x2))
			}
		}
	}

	t.Log("Test Count:", count)
}
