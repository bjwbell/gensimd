//

package main

import (
	"fmt"
	"math/rand"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -fn "regspill1" -outfn "regspill1" -f "reg_spill1_simd.go" -o "reg_spill1_amd64.s" -goprotofile "reg_spill1_simd_proto.go"

func regspillsRef(x, y int32) int32 {
	dist := int32(0)
	xi := x
	xj := x
	yi := y
	yj := y
	dx := xj + xi
	dy := yj + yi
	sqX := dx * dx
	sqY := dy * dy
	sqDist := sqX + sqY
	dx2 := dx - dy
	t := 2*dx2 + dy
	dist = sqDist + t
	return dist
}

func main() {
	errors := 0

	n := 1024
	x := make([]simd.I32x4, n)
	y := make([]simd.I32x4, n)
	rand.Seed(42)
	maxCoord := int32(255)
	for i := 0; i < n; i++ {
		x[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
		y[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
	}
	val := regspill1(x[0][0], y[0][0])
	if val != regspillsRef(x[0][0], y[0][0]) {
		fmt.Printf("SIMD %v != reference (%v)\n", val, regspillsRef(x[0][0], y[0][0]))
		errors++
	}

	fmt.Printf("Done, errors: %d\n", errors)
}
