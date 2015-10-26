package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/bjwbell/gensimd/simd"
)

func main() {

	count := 0
	rand.Seed(42)
	errors := 0
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

			shift := uint8(j)

			var xI32x4 simd.I32x4
			var yI32x4 simd.I32x4

			var xF32x4 simd.F32x4
			var yF32x4 simd.F32x4

			for idx := 0; idx < 4; idx++ {
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

				xI32x4[idx] = int32(rand.Intn(abs_a))
				yI32x4[idx] = int32(rand.Intn(abs_b))
				xF32x4[idx] = rand.Float32()
				yF32x4[idx] = rand.Float32()
			}

			if addi32x4(xI32x4, yI32x4) != simd.AddI32x4(xI32x4, yI32x4) {
				fmt.Printf("addi32x4(%v, %v)\n", xI32x4, yI32x4)
				fmt.Println("x:", xI32x4)
				fmt.Println("y:", yI32x4)
				fmt.Println("s(xI32x4, yI32x4):", addi32x4(xI32x4, yI32x4))
				fmt.Println(" (xI32x4, yI32x4):", simd.AddI32x4(xI32x4, yI32x4))
				errors++
			}
			if subi32x4(xI32x4, yI32x4) != simd.SubI32x4(xI32x4, yI32x4) {
				fmt.Printf("subi32x4(%v, %v)\n", xI32x4, yI32x4)
				fmt.Println("x:", xI32x4)
				fmt.Println("y:", yI32x4)
				fmt.Println("s:", subi32x4(xI32x4, yI32x4))
				fmt.Println(" :", simd.SubI32x4(xI32x4, yI32x4))
				errors++
			}
			if muli32x4(xI32x4, yI32x4) != simd.MulI32x4(xI32x4, yI32x4) {
				fmt.Printf("muli32x4(%v, %v)\n", xI32x4, yI32x4)
				fmt.Println("x:", xI32x4)
				fmt.Println("y:", yI32x4)
				fmt.Println("s:", muli32x4(xI32x4, yI32x4))
				fmt.Println(" :", simd.MulI32x4(xI32x4, yI32x4))
				errors++
			}
			if shli32x4(xI32x4, shift) != simd.ShlI32x4(xI32x4, shift) {
				fmt.Printf("shli32x4(%v, %v)\n", xI32x4, shift)
				fmt.Println("x:", xI32x4)
				fmt.Println("shift:", shift)
				fmt.Println("s:", shli32x4(xI32x4, shift))
				fmt.Println(" :", simd.ShlI32x4(xI32x4, shift))
				errors++
			}
			if shri32x4(xI32x4, shift) != simd.ShrI32x4(xI32x4, shift) {
				fmt.Printf("shri32x4(%v, %v)\n", xI32x4, shift)
				fmt.Println("x:", xI32x4)
				fmt.Println("shift:", shift)
				fmt.Println("s:", shri32x4(xI32x4, shift))
				fmt.Println(" :", simd.ShrI32x4(xI32x4, shift))
				errors++
			}

			if addf32x4(xF32x4, yF32x4) != simd.AddF32x4(xF32x4, yF32x4) {
				fmt.Printf("addf32x4(%v, %v)\n", xF32x4, yF32x4)
				fmt.Println("x:", xF32x4)
				fmt.Println("y:", yF32x4)
				fmt.Println("s:", addf32x4(xF32x4, yF32x4))
				fmt.Println(" :", simd.AddF32x4(xF32x4, yF32x4))
				errors++
			}
			if subf32x4(xF32x4, yF32x4) != simd.SubF32x4(xF32x4, yF32x4) {
				fmt.Printf("subf32x4(%v, %v)\n", xF32x4, yF32x4)
				fmt.Println("x:", xF32x4)
				fmt.Println("y:", yF32x4)
				fmt.Println("s:", subf32x4(xF32x4, yF32x4))
				fmt.Println(" :", simd.SubF32x4(xF32x4, yF32x4))
				errors++
			}
			if mulf32x4(xF32x4, yF32x4) != simd.MulF32x4(xF32x4, yF32x4) {
				fmt.Printf("mulf32x4(%v, %v)\n", xF32x4, yF32x4)
				fmt.Println("x:", xF32x4)
				fmt.Println("y:", yF32x4)
				fmt.Println("s:", mulf32x4(xF32x4, yF32x4))
				fmt.Println(" :", simd.MulF32x4(xF32x4, yF32x4))
			}
			if divf32x4(xF32x4, yF32x4) != simd.DivF32x4(xF32x4, yF32x4) {
				fmt.Printf("divf32x4(%v, %v)\n", xF32x4, yF32x4)
				fmt.Println("x:", xF32x4)
				fmt.Println("y:", yF32x4)
				fmt.Println("s:", divf32x4(xF32x4, yF32x4))
				fmt.Println(" :", simd.DivF32x4(xF32x4, yF32x4))
				errors++
			}

		}
	}

	fmt.Printf("Done, errors: %d\n", errors)
}
