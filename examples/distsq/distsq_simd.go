// +build !amd64,gc

package main

import (
	"math"

	"github.com/bjwbell/gensimd/simd"
)

func distsq(x, y []simd.I32x4) int32 {
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
			dx := simd.SubI32x4(x[j], x[i])
			dy := simd.SubI32x4(y[j], y[i])
			sqX := simd.MulI32x4(dx, dx)
			sqY := simd.MulI32x4(dy, dy)
			sqDist := simd.AddI32x4(sqX, sqY)
			if sqDist[0] < dist {
				dist = sqDist[0]
			}
			if sqDist[1] < dist {
				dist = sqDist[1]
			}
			if sqDist[2] < dist {
				dist = sqDist[2]
			}
			if sqDist[3] < dist {
				dist = sqDist[3]
			}

		}
	}
	return dist
}
