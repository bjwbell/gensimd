// +build amd64,gc

package tests

import (
	"math"
	"math/rand"
	"testing"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -fn "regression1Simd" -outfn "regression1Simds" -f "regression1_simd_test.go" -o "regression1_simd_test_amd64.s"

func regression1Simds(x []simd.I32x4, y []simd.I32x4) int32

func regression1Simd(x, y []simd.I32x4) int32 {
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
	dy2 := simd.SubI32x4(sqX, sqY)
	dist = sqDist[0] + dy2[2]
	return dist
}

func TestRegression1RegSpill(t *testing.T) {
	errors := 0
	count := 1
	n := 1024
	x := make([]simd.I32x4, n)
	y := make([]simd.I32x4, n)
	rand.Seed(42)
	maxCoord := int32(255)
	for i := 0; i < n; i++ {
		x[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
		y[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
	}
	val := regression1Simds(x, y)
	if val != regression1Simd(x, y) {
		t.Errorf("SIMD %v != reference (%v)\n", val, regression1Simd(x, y))
		errors++
	}

	t.Log("Test Count:", count)
}
