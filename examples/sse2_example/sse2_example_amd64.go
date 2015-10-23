// +build amd64,gc

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -debug -fn "addpd" -outfn "addpd" -f "sse2_example_other.go" -o "sse2_example_amd64.s"

func addpd(x, y simd.M128d) simd.M128d

func main() {
	// bail if sse2 is not supported
	if !simd.SSE2() {
		fmt.Println("SSE2 not available")
		return
	}

	count := 0
	errors := 0
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
				fmt.Printf("addm128d(%v, %v)\n", xF64x2, yF64x2)
				fmt.Println("x:", xF64x2)
				fmt.Println("y:", yF64x2)
				fmt.Println("s:", addpd(xM128d, yM128d))
				fmt.Println(" :", simd.AddF64x2(xF64x2, yF64x2))
				errors++
			}
		}
	}

	fmt.Printf("Done checking sse2 example results, errors: %d\n", errors)
}
