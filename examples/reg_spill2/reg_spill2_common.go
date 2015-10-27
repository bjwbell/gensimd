//

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -spills -debug -fn "regspill2" -outfn "regspill2" -f "reg_spill2_simd.go" -o "reg_spill2_amd64.s" -goprotofile "reg_spill2_simd_proto.go"

func regspillsRef(x, y []simd.I32x4) int32 {
	if len(x) != len(y) {
		return -1
	}
	dist := int32(math.MaxInt32)
	i := 0
	j := 1
	dx := simd.SubI32x4(x[j], x[i])
	dy := simd.SubI32x4(y[j], y[i])
	sqX := simd.MulI32x4(dx, dx)
	sqY := simd.MulI32x4(dy, dy)
	sqDist := simd.AddI32x4(sqX, sqY)
	dx2 := simd.SubI32x4(dx, dy)

	dy2 := simd.SubI32x4(sqX, sqY)
	sqX2 := simd.MulI32x4(dx2, dx2)
	sqY2 := simd.MulI32x4(dy2, dy2)
	sqDist2 := simd.AddI32x4(sqX2, sqY2)
	dx3 := simd.SubI32x4(dx2, dy2)
	dy3 := simd.SubI32x4(sqX2, sqY2)
	sqX3 := simd.MulI32x4(dx3, dx3)
	sqY3 := simd.MulI32x4(dy3, dy3)
	sqDist3 := simd.AddI32x4(sqX3, sqY3)
	t := dx[0] + dx[1] + dx[2] + dx[3] +
		dx2[0] + dx2[1] + dx2[2] + dx2[3] +
		dx3[0] + dx3[1] + dx3[2] + dx3[3]
	v := dy[0] + dy[1] + dy[2] + dy[3] +
		dy2[0] + dy2[1] + dy2[2] + dy2[3] +
		dy3[0] + dy3[1] + dy3[2] + dy3[3]
	dist = t + v + sqDist[0] + sqDist2[1] + sqDist3[2]
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
	val := regspill2(x, y)
	if val != regspillsRef(x, y) {
		fmt.Printf("SIMD %v != reference (%v)\n", val, regspillsRef(x, y))
		errors++
	}

	fmt.Printf("Done, errors: %d\n", errors)
}
