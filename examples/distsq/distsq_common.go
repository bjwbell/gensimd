//

package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -fn "distsq" -outfn "distsq" -f "distsq_simd.go" -o "distsq_amd64.s" -goprotofile "distsq_simd_proto.go"

func distsqRef(x, y []simd.I32x4) int32 {
	if len(x) != len(y) {
		return -1
	}
	n := len(x)
	dist := int32(math.MaxInt32)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for idx := 0; idx < 4; idx++ {
				dx := x[j][idx] - x[i][idx]
				dy := y[j][idx] - y[i][idx]
				sqX := dx * dx
				sqY := dy * dy
				sqDist := sqX + sqY
				if sqDist < dist {
					dist = sqDist
				}
			}
		}
	}
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
	distSq := distsq(x, y)
	if distSq != distsqRef(x, y) {
		fmt.Printf("SIMD %v != reference (%v)\n", distSq, distsqRef(x, y))
		errors++
	}

	fmt.Printf("Done, errors: %d\n", errors)
}
